package server

import (
	"context"
	"fmt"
	globallog "github.com/seoyhaein/caleb/log"
	"github.com/seoyhaein/caleb/v1rpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const (
	defaultMaxRequestBytes   = 1.5 * 1024 * 1024
	defaultGrpcOverheadBytes = 512 * 1024
	defaultMaxStreams        = 1<<32 - 1 // math.MaxUint32와 동일
	defaultMaxSendBytes      = 1<<31 - 1 // math.MaxInt32와 동일
)

var (
	Address = ":50052" // default value
	logger  = globallog.Log
)

func init() {
	// TODO: Prometheus 적용 예정
}

// 값이 없거나 잘못된 경우 defaultVal 을 반환한다.
func getEnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		logger.Infof("Invalid value for %s: %v. Using default: %d", key, err, defaultVal)
		return defaultVal
	}
	return val
}

// gRPC 요청을 받을 때마다 요청 메서드와 에러 정보를 로깅함.
func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger.Infof("Received request for %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		logger.Infof("Method %s error: %v", info.FullMethod, err)
	}
	return resp, err
}

func Server() error {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", Address, err)
	}

	// TLS 설정 예시 (TLS가 필요하면 주석 풀고 사용)
	/*
		certFile := "server.crt"
		keyFile := "server.key"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return fmt.Errorf("failed to load TLS credentials: %w", err)
		}
	*/

	// 환경 변수로 옵션 값을 오버라이드할 수 있음
	maxRecvMsgSize := getEnvInt("GRPC_MAX_RECV_MSG_SIZE", int(defaultMaxRequestBytes+defaultGrpcOverheadBytes))
	maxSendMsgSize := getEnvInt("GRPC_MAX_SEND_MSG_SIZE", defaultMaxSendBytes)
	maxConcurrentStreams := getEnvInt("GRPC_MAX_CONCURRENT_STREAMS", defaultMaxStreams)

	// gRPC 서버 옵션 설정
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxRecvMsgSize),
		grpc.MaxSendMsgSize(maxSendMsgSize),
		grpc.MaxConcurrentStreams(uint32(maxConcurrentStreams)),
		grpc.UnaryInterceptor(loggingInterceptor),
		// grpc.Creds(creds), // TLS 사용 시 활성화
	}

	grpcServer := grpc.NewServer(opts...)

	// 기존 서비스 등록 (주석 처리된 상태)
	// 수정: 헬스 체크 서비스 등록
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// 기존: Reflection 서비스 등록, 디버깅 및 grpcurl 노출 위해서.
	reflection.Register(grpcServer)
	service.RegisterStageBuilderServiceServer(grpcServer)
	logger.Infof("gRPC server started, address: %s", Address)

	// graceful shutdown 처리 추가
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigCh
		logger.Infof("Received signal: %v. Initiating graceful shutdown...", sig)
		// GracefulStop 은 현재 처리 중인 요청을 모두 완료한 후 서버를 중지함.
		grpcServer.GracefulStop()
	}()

	// 서버 시작
	serveErr := grpcServer.Serve(lis)
	if serveErr != nil {
		if !strings.Contains(serveErr.Error(), "use of closed network connection") {
			logger.Infof("gRPC server returned with error: %v", serveErr)
		} else {
			logger.Infof("gRPC server is shut down")
		}
	}
	return serveErr
}

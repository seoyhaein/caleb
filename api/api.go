package api

import (
	"context"
	pb "github.com/seoyhaein/caleb/protos"
)

// PipeApis dockerfile 검증 관련된 인터페이스 grpc 연결 및 cli 연결 목적
type PipeApis interface {
	BuildStageImage(context.Context, *pb.DockerfileRequest) (*pb.DockerfileResponse, error)
}

// pipeApisImpl PipeApis 인터페이스의 구현체
type pipeApisImpl struct{}

// NewPipeApis PipeApis 인터페이스의 구현체를 생성하는 factory 함수
func NewPipeApis() PipeApis {
	return &pipeApisImpl{}
}

func (f *pipeApisImpl) BuildStageImage(context.Context, *pb.DockerfileRequest) (*pb.DockerfileResponse, error) {
	return nil, nil
}

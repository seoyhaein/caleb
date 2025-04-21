package service

import (
	"github.com/seoyhaein/caleb/api"
	pb "github.com/seoyhaein/caleb/protos"
	"google.golang.org/grpc"
)

type stageBuilderServiceServerImpl struct {
	pb.UnimplementedStageBuilderServiceServer
	pipeApis api.PipeApis
}

func NewStageBuilderServiceServer() pb.StageBuilderServiceServer {
	return &stageBuilderServiceServerImpl{
		pipeApis: api.NewPipeApis(),
	}
}

func RegisterStageBuilderServiceServer(service *grpc.Server) {
	pb.RegisterStageBuilderServiceServer(service, NewStageBuilderServiceServer())
}

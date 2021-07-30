package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pbServer "github.com/herryg91/dply/dply-server/handler/grst/server"
)

type handlerServer struct {
	pbServer.UnimplementedServerApiServer
}

func NewServerHandler() pbServer.ServerApiServer {
	return &handlerServer{}
}
func (h *handlerServer) Status(ctx context.Context, req *empty.Empty) (*pbServer.StatusResp, error) {
	return &pbServer.StatusResp{
		Status: "ok",
	}, nil
}

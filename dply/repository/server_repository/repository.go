package server_repository

import (
	"context"

	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbServer "github.com/herryg91/dply/dply/clients/grst/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository struct {
	cli pbServer.ServerApiClient
}

func New(cli pbServer.ServerApiClient) repository_intf.ServerRepository {
	return &repository{cli}
}

func (r *repository) Status() bool {
	resp, err := r.cli.Status(context.Background(), &emptypb.Empty{})
	if err != nil {
		return false
	}

	return resp.Status == "ok"
}

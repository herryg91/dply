package handler

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"github.com/herryg91/dply/dply-server/entity"
	pbSpec "github.com/herryg91/dply/dply-server/handler/grst/spec"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"google.golang.org/grpc/codes"
)

func (h *specConfig) GetScale(ctx context.Context, req *pbSpec.GetScaleReq) (*pbSpec.Scale, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.scale_uc.Get(req.Env, req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 12001, err.Error())
	}

	return &pbSpec.Scale{
		Env:                  resp.Env,
		Name:                 resp.Name,
		MinReplica:           int32(resp.MinReplica),
		MaxReplica:           int32(resp.MaxReplica),
		MinCpu:               int32(resp.MinCpu),
		MaxCpu:               int32(resp.MaxCpu),
		MinMemory:            int32(resp.MinMemory),
		MaxMemory:            int32(resp.MaxMemory),
		TargetCPUUtilization: int32(resp.TargetCPUUtilization),
	}, nil
}

func (h *specConfig) UpsertScale(ctx context.Context, req *pbSpec.UpsertScaleReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	err := h.scale_uc.Upsert(entity.Scale{
		Env:                  req.Env,
		Name:                 req.Name,
		MinReplica:           int(req.MinReplica),
		MaxReplica:           int(req.MaxReplica),
		MinCpu:               int(req.MinCpu),
		MaxCpu:               int(req.MaxCpu),
		MinMemory:            int(req.MinMemory),
		MaxMemory:            int(req.MaxMemory),
		TargetCPUUtilization: int(req.TargetCPUUtilization),
		CreatedBy:            userCtx.Id,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 12101, err.Error())
	}

	return &empty.Empty{}, nil
}

package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"github.com/herryg91/dply/dply-server/entity"
	pbSpec "github.com/herryg91/dply/dply-server/handler/grst/spec"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"google.golang.org/grpc/codes"
)

func (h *specConfig) GetEnvar(ctx context.Context, req *pbSpec.GetEnvarReq) (*pbSpec.Envar, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}

	resp, err := h.envar_uc.Get(req.Project, req.Env, req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 11001, err.Error())
	}

	variables, _ := json.Marshal(&resp.Variables)

	return &pbSpec.Envar{
		Variables: string(variables),
	}, nil
}

func (h *specConfig) UpsertEnvar(ctx context.Context, req *pbSpec.UpsertEnvarReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	variables := map[string]interface{}{}
	json.Unmarshal([]byte(req.Variables), &variables)
	err := h.envar_uc.Upsert(entity.Envar{
		Project:   req.Project,
		Env:       req.Env,
		Name:      req.Name,
		Variables: variables,
		CreatedBy: userCtx.Id,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 11101, err.Error())
	}

	return &empty.Empty{}, nil
}

package handler

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	deploy_usecase "github.com/herryg91/dply/dply-server/app/usecase/deploy"
	pbDeploy "github.com/herryg91/dply/dply-server/handler/grst/deploy"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"google.golang.org/grpc/codes"
)

type handlerDeploy struct {
	deploy_uc deploy_usecase.UseCase
	pbDeploy.UnimplementedDeployApiServer
}

func NewDeployHandler(deploy_uc deploy_usecase.UseCase) pbDeploy.DeployApiServer {
	return &handlerDeploy{deploy_uc: deploy_uc}
}

func (h *handlerDeploy) DeployImage(ctx context.Context, req *pbDeploy.DeployImageReq) (*empty.Empty, error) {
	if err := pbDeploy.ValidateRequest(req); err != nil {
		return nil, err
	}

	userCtx := interceptor.ExtractMustLoginContext(ctx)

	err := h.deploy_uc.DeployImage(req.Project, req.Env, req.Name, req.Digest, userCtx.Id)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 11001, err.Error())
	}

	return &empty.Empty{}, nil
}
func (h *handlerDeploy) Redeploy(ctx context.Context, req *pbDeploy.RedeployReq) (*empty.Empty, error) {
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	err := h.deploy_uc.Redeploy(req.Project, req.Env, req.Name, userCtx.Id)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 12001, err.Error())
	}

	return &empty.Empty{}, nil
}

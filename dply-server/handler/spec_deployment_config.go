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

func (h *specConfig) GetDeploymentConfig(ctx context.Context, req *pbSpec.GetDeploymentConfigReq) (*pbSpec.DeploymentConfig, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.deployment_config_uc.Get(req.Project, req.Env, req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14001, err.Error())
	}

	var liveness_probe, readiness_probe, startup_probe *pbSpec.HttpGetProbe
	if resp.LivenessProbe != nil {
		liveness_probe = &pbSpec.HttpGetProbe{
			Path:                resp.LivenessProbe.Path,
			Port:                int32(resp.LivenessProbe.Port),
			FailureThreshold:    int32(resp.LivenessProbe.FailureThreshold),
			PeriodSeconds:       int32(resp.LivenessProbe.PeriodSeconds),
			InitialDelaySeconds: int32(resp.LivenessProbe.InitialDelaySeconds),
		}
	} else {
		liveness_probe = nil
	}

	if resp.ReadinessProbe != nil {
		readiness_probe = &pbSpec.HttpGetProbe{
			Path:                resp.ReadinessProbe.Path,
			Port:                int32(resp.ReadinessProbe.Port),
			FailureThreshold:    int32(resp.ReadinessProbe.FailureThreshold),
			PeriodSeconds:       int32(resp.ReadinessProbe.PeriodSeconds),
			InitialDelaySeconds: int32(resp.ReadinessProbe.InitialDelaySeconds),
		}
	} else {
		readiness_probe = nil
	}

	if resp.StartupProbe != nil {
		startup_probe = &pbSpec.HttpGetProbe{
			Path:                resp.StartupProbe.Path,
			Port:                int32(resp.StartupProbe.Port),
			FailureThreshold:    int32(resp.StartupProbe.FailureThreshold),
			PeriodSeconds:       int32(resp.StartupProbe.PeriodSeconds),
			InitialDelaySeconds: int32(resp.StartupProbe.InitialDelaySeconds),
		}
	} else {
		startup_probe = nil
	}

	return &pbSpec.DeploymentConfig{
		LivenessProbe:  liveness_probe,
		ReadinessProbe: readiness_probe,
		StartupProbe:   startup_probe,
	}, nil
}

func (h *specConfig) UpsertDeploymentConfig(ctx context.Context, req *pbSpec.UpsertDeploymentConfigReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	var liveness_probe, readiness_probe, startup_probe *entity.HttpGetProbe
	if req.LivenessProbe != nil {
		liveness_probe = &entity.HttpGetProbe{
			Path:                req.LivenessProbe.Path,
			Port:                int(req.LivenessProbe.Port),
			FailureThreshold:    int(req.LivenessProbe.FailureThreshold),
			PeriodSeconds:       int(req.LivenessProbe.PeriodSeconds),
			InitialDelaySeconds: int(req.LivenessProbe.InitialDelaySeconds),
		}
	} else {
		liveness_probe = nil
	}
	if req.ReadinessProbe != nil {
		readiness_probe = &entity.HttpGetProbe{
			Path:                req.ReadinessProbe.Path,
			Port:                int(req.ReadinessProbe.Port),
			FailureThreshold:    int(req.ReadinessProbe.FailureThreshold),
			PeriodSeconds:       int(req.ReadinessProbe.PeriodSeconds),
			InitialDelaySeconds: int(req.ReadinessProbe.InitialDelaySeconds),
		}
	} else {
		readiness_probe = nil
	}

	if req.StartupProbe != nil {
		startup_probe = &entity.HttpGetProbe{
			Path:                req.StartupProbe.Path,
			Port:                int(req.StartupProbe.Port),
			FailureThreshold:    int(req.StartupProbe.FailureThreshold),
			PeriodSeconds:       int(req.StartupProbe.PeriodSeconds),
			InitialDelaySeconds: int(req.StartupProbe.InitialDelaySeconds),
		}
	} else {
		startup_probe = nil
	}

	err := h.deployment_config_uc.Upsert(entity.DeploymentConfig{
		Project:        req.Project,
		Env:            req.Env,
		Name:           req.Name,
		LivenessProbe:  liveness_probe,
		ReadinessProbe: readiness_probe,
		StartupProbe:   startup_probe,
		CreatedBy:      userCtx.Id,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14101, err.Error())
	}

	return &empty.Empty{}, nil
}

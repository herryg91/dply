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

func (h *specConfig) GetPort(ctx context.Context, req *pbSpec.GetPortReq) (*pbSpec.Ports, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.port_uc.Get(req.Env, req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 13001, err.Error())
	}

	ports := []*pbSpec.Port{}
	for _, p := range resp.Ports {
		ports = append(ports, &pbSpec.Port{
			PortName: p.Name,
			Port:     int32(p.Port),
			Protocol: string(p.Protocol),
		})
	}

	return &pbSpec.Ports{
		Ports: ports,
	}, nil
}
func (h *specConfig) UpsertPort(ctx context.Context, req *pbSpec.UpsertPortReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	ports := []entity.PortSpec{}
	for _, p := range req.Ports {
		ports = append(ports, entity.PortSpec{
			Name:     p.PortName,
			Port:     int(p.Port),
			Protocol: entity.PortType(p.Protocol),
		})
	}

	err := h.port_uc.Upsert(entity.Port{
		Env:       req.Env,
		Name:      req.Name,
		Ports:     ports,
		CreatedBy: userCtx.Id,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 13101, err.Error())
	}

	return &empty.Empty{}, nil
}
func (h *specConfig) GetPortTemplate(ctx context.Context, req *pbSpec.GetPortTemplateReq) (*pbSpec.Ports, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.port_uc.GetTemplate(req.TemplateName)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 13201, err.Error())
	}

	ports := []*pbSpec.Port{}
	for _, p := range resp.Ports {
		ports = append(ports, &pbSpec.Port{
			PortName: p.Name,
			Port:     int32(p.Port),
			Protocol: string(p.Protocol),
		})
	}

	return &pbSpec.Ports{
		Ports: ports,
	}, nil
}
func (h *specConfig) UpdatePortTemplate(ctx context.Context, req *pbSpec.UpdatePortTemplateReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}

	ports := []entity.PortSpec{}
	for _, p := range req.Ports {
		ports = append(ports, entity.PortSpec{
			Name:     p.PortName,
			Port:     int(p.Port),
			Protocol: entity.PortType(p.Protocol),
		})
	}

	err := h.port_uc.UpsertTemplate(entity.PortTemplate{
		TemplateName: req.TemplateName,
		Ports:        ports,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 13401, err.Error())
	}

	return &empty.Empty{}, nil
}

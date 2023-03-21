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

func (h *specConfig) GetAffinity(ctx context.Context, req *pbSpec.GetAffinityReq) (*pbSpec.Affinity, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.affinity_uc.Get(req.Project, req.Env, req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14001, err.Error())
	}

	nodeAffinity := []*pbSpec.AffinityTerm{}
	podAffinity := []*pbSpec.AffinityTerm{}
	podAntiAffinity := []*pbSpec.AffinityTerm{}
	tolerations := []*pbSpec.AffinityToleration{}

	for _, a := range resp.NodeAffinity {
		nodeAffinity = append(nodeAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAffinity {
		podAffinity = append(podAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.Tolerations {
		tolerations = append(tolerations, &pbSpec.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	return &pbSpec.Affinity{
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
		Tolerations:     tolerations,
	}, nil
}

func (h *specConfig) UpsertAffinity(ctx context.Context, req *pbSpec.UpsertAffinityReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	userCtx := interceptor.ExtractMustLoginContext(ctx)

	nodeAffinity := []entity.AffinityTerm{}
	podAffinity := []entity.AffinityTerm{}
	podAntiAffinity := []entity.AffinityTerm{}
	tolerations := []entity.AffinityToleration{}

	for _, a := range req.NodeAffinity {
		nodeAffinity = append(nodeAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.PodAffinity {
		podAffinity = append(podAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.Tolerations {
		tolerations = append(tolerations, entity.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	err := h.affinity_uc.Upsert(entity.Affinity{
		Project:         req.Project,
		Env:             req.Env,
		Name:            req.Name,
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
		Tolerations:     tolerations,
		CreatedBy:       userCtx.Id,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14101, err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *specConfig) GetAffinityTemplate(ctx context.Context, req *pbSpec.GetAffinityTemplateReq) (*pbSpec.Affinity, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}
	resp, err := h.affinity_uc.GetTemplate(req.TemplateName)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14201, err.Error())
	}

	nodeAffinity := []*pbSpec.AffinityTerm{}
	podAffinity := []*pbSpec.AffinityTerm{}
	podAntiAffinity := []*pbSpec.AffinityTerm{}
	tolerations := []*pbSpec.AffinityToleration{}

	for _, a := range resp.NodeAffinity {
		nodeAffinity = append(nodeAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAffinity {
		podAffinity = append(podAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.Tolerations {
		tolerations = append(tolerations, &pbSpec.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	return &pbSpec.Affinity{
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
		Tolerations:     tolerations,
	}, nil
}

func (h *specConfig) UpdateAffinityTemplate(ctx context.Context, req *pbSpec.UpdateAffinityTemplateReq) (*empty.Empty, error) {
	if err := pbSpec.ValidateRequest(req); err != nil {
		return nil, err
	}

	nodeAffinity := []entity.AffinityTerm{}
	podAffinity := []entity.AffinityTerm{}
	podAntiAffinity := []entity.AffinityTerm{}
	tolerations := []entity.AffinityToleration{}

	for _, a := range req.NodeAffinity {
		nodeAffinity = append(nodeAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.PodAffinity {
		podAffinity = append(podAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range req.Tolerations {
		tolerations = append(tolerations, entity.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	err := h.affinity_uc.UpsertTemplate(entity.AffinityTemplate{
		TemplateName:    req.TemplateName,
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
		Tolerations:     tolerations,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14401, err.Error())
	}

	return &empty.Empty{}, nil
}

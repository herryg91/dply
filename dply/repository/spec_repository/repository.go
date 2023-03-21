package spec_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbSpec "github.com/herryg91/dply/dply/clients/grst/spec"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
)

type repository struct {
	cli pbSpec.SpecApiClient
}

func New(cli pbSpec.SpecApiClient) repository_intf.SpecRepository {
	return &repository{cli}
}
func (r *repository) GetEnvar(project, env, name string) (*entity.Envar, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetEnvar(ctx, &pbSpec.GetEnvarReq{Project: project, Env: env, Name: name})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	envar := &entity.Envar{Project: project, Env: env, Name: name, Variables: map[string]interface{}{}}
	json.Unmarshal([]byte(resp.Variables), &envar.Variables)

	return envar, nil
}
func (r *repository) UpsertEnvar(data entity.Envar) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	toUpperVariable := map[string]interface{}{}
	for k, v := range data.Variables {
		toUpperVariable[strings.ToUpper(k)] = v
	}

	variables, _ := json.Marshal(&toUpperVariable)
	_, err := r.cli.UpsertEnvar(ctx, &pbSpec.UpsertEnvarReq{
		Project:   data.Project,
		Env:       data.Env,
		Name:      data.Name,
		Variables: string(variables),
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) GetScale(project, env, name string) (*entity.Scale, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetScale(ctx, &pbSpec.GetScaleReq{Project: project, Env: env, Name: name})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	scale := &entity.Scale{
		Project:              project,
		Env:                  env,
		Name:                 name,
		MinReplica:           int(resp.MinReplica),
		MaxReplica:           int(resp.MaxReplica),
		MinCpu:               int(resp.MinCpu),
		MaxCpu:               int(resp.MaxCpu),
		MinMemory:            int(resp.MinMemory),
		MaxMemory:            int(resp.MaxMemory),
		TargetCPUUtilization: int(resp.TargetCPUUtilization),
	}

	return scale, nil
}
func (r *repository) UpsertScale(data entity.Scale) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.UpsertScale(ctx, &pbSpec.UpsertScaleReq{
		Project:              data.Project,
		Env:                  data.Env,
		Name:                 data.Name,
		MinReplica:           int32(data.MinReplica),
		MaxReplica:           int32(data.MaxReplica),
		MinCpu:               int32(data.MinCpu),
		MaxCpu:               int32(data.MaxCpu),
		MinMemory:            int32(data.MinMemory),
		MaxMemory:            int32(data.MaxMemory),
		TargetCPUUtilization: int32(data.TargetCPUUtilization),
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) GetPort(project, env, name string) (*entity.Port, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetPort(ctx, &pbSpec.GetPortReq{Project: project, Env: env, Name: name})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	portSpecs := []entity.PortSpec{}
	for _, p := range resp.Ports {
		portSpecs = append(portSpecs, entity.PortSpec{
			Name:       p.PortName,
			Port:       int(p.Port),
			RemotePort: int(p.RemotePort),
			Protocol:   entity.PortType(p.Protocol),
		})
	}
	port := &entity.Port{Project: project, Env: env, Name: name, AccessType: entity.AccessType(resp.AccessType), ExternalIP: resp.ExternalIP, Ports: portSpecs}
	return port, nil
}
func (r *repository) UpsertPort(data entity.Port) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	portParam := []*pbSpec.Port{}
	for _, p := range data.Ports {
		portParam = append(portParam, &pbSpec.Port{
			PortName:   p.Name,
			Port:       int32(p.Port),
			RemotePort: int32(p.RemotePort),
			Protocol:   string(p.Protocol),
		})
	}
	_, err := r.cli.UpsertPort(ctx, &pbSpec.UpsertPortReq{
		Project:    data.Project,
		Env:        data.Env,
		Name:       data.Name,
		AccessType: string(data.AccessType),
		ExternalIP: data.ExternalIP,
		Ports:      portParam,
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) GetAffinity(project, env, name string) (*entity.Affinity, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetAffinity(ctx, &pbSpec.GetAffinityReq{Project: project, Env: env, Name: name})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	nodeAffinity := []entity.AffinityTerm{}
	podAffinity := []entity.AffinityTerm{}
	podAntiAffinity := []entity.AffinityTerm{}
	tolerations := []entity.AffinityToleration{}

	for _, a := range resp.NodeAffinity {
		nodeAffinity = append(nodeAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAffinity {
		podAffinity = append(podAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.Tolerations {
		tolerations = append(tolerations, entity.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	affinity := &entity.Affinity{Project: project, Env: env, Name: name, NodeAffinity: nodeAffinity, PodAffinity: podAffinity, PodAntiAffinity: podAntiAffinity, Tolerations: tolerations}

	return affinity, nil
}
func (r *repository) UpsertAffinity(data entity.Affinity) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	nodeAffinity := []*pbSpec.AffinityTerm{}
	podAffinity := []*pbSpec.AffinityTerm{}
	podAntiAffinity := []*pbSpec.AffinityTerm{}
	tolerations := []*pbSpec.AffinityToleration{}

	for _, a := range data.NodeAffinity {
		nodeAffinity = append(nodeAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}

	for _, a := range data.PodAffinity {
		podAffinity = append(podAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range data.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range data.Tolerations {
		tolerations = append(tolerations, &pbSpec.AffinityToleration{Key: a.Key, Operator: a.Operator, Value: a.Value, Effect: a.Effect})
	}

	_, err := r.cli.UpsertAffinity(ctx, &pbSpec.UpsertAffinityReq{
		Project:         data.Project,
		Env:             data.Env,
		Name:            data.Name,
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
		Tolerations:     tolerations,
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) GetDeploymentConfig(project, env, name string) (*entity.DeploymentConfig, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetDeploymentConfig(ctx, &pbSpec.GetDeploymentConfigReq{Project: project, Env: env, Name: name})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	out := &entity.DeploymentConfig{Project: project, Env: env, Name: name,
		LivenessProbe:  nil,
		ReadinessProbe: nil,
		StartupProbe:   nil,
	}

	if resp.LivenessProbe != nil && resp.LivenessProbe.Path != "" && resp.LivenessProbe.Port > 0 {
		out.LivenessProbe = &entity.HttpGetProbe{
			Path:                resp.LivenessProbe.Path,
			Port:                int(resp.LivenessProbe.Port),
			FailureThreshold:    int(resp.LivenessProbe.FailureThreshold),
			PeriodSeconds:       int(resp.LivenessProbe.PeriodSeconds),
			InitialDelaySeconds: int(resp.LivenessProbe.InitialDelaySeconds),
		}
	}

	if resp.ReadinessProbe != nil && resp.ReadinessProbe.Path != "" && resp.ReadinessProbe.Port > 0 {
		out.ReadinessProbe = &entity.HttpGetProbe{
			Path:                resp.ReadinessProbe.Path,
			Port:                int(resp.ReadinessProbe.Port),
			FailureThreshold:    int(resp.ReadinessProbe.FailureThreshold),
			PeriodSeconds:       int(resp.ReadinessProbe.PeriodSeconds),
			InitialDelaySeconds: int(resp.ReadinessProbe.InitialDelaySeconds),
		}
	}

	if resp.StartupProbe != nil && resp.StartupProbe.Path != "" && resp.StartupProbe.Port > 0 {
		out.StartupProbe = &entity.HttpGetProbe{
			Path:                resp.StartupProbe.Path,
			Port:                int(resp.StartupProbe.Port),
			FailureThreshold:    int(resp.StartupProbe.FailureThreshold),
			PeriodSeconds:       int(resp.StartupProbe.PeriodSeconds),
			InitialDelaySeconds: int(resp.StartupProbe.InitialDelaySeconds),
		}
	}
	return out, nil
}

func (r *repository) UpsertDeploymentConfig(data entity.DeploymentConfig) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	param := &pbSpec.UpsertDeploymentConfigReq{
		Project:        data.Project,
		Env:            data.Env,
		Name:           data.Name,
		LivenessProbe:  nil,
		ReadinessProbe: nil,
		StartupProbe:   nil,
	}
	if data.LivenessProbe != nil && data.LivenessProbe.Path != "" && data.LivenessProbe.Port > 0 {
		param.LivenessProbe = &pbSpec.HttpGetProbe{
			Path:                data.LivenessProbe.Path,
			Port:                int32(data.LivenessProbe.Port),
			FailureThreshold:    int32(data.LivenessProbe.FailureThreshold),
			PeriodSeconds:       int32(data.LivenessProbe.PeriodSeconds),
			InitialDelaySeconds: int32(data.LivenessProbe.InitialDelaySeconds),
		}
	}
	if data.ReadinessProbe != nil && data.ReadinessProbe.Path != "" && data.ReadinessProbe.Port > 0 {
		param.ReadinessProbe = &pbSpec.HttpGetProbe{
			Path:                data.ReadinessProbe.Path,
			Port:                int32(data.ReadinessProbe.Port),
			FailureThreshold:    int32(data.ReadinessProbe.FailureThreshold),
			PeriodSeconds:       int32(data.ReadinessProbe.PeriodSeconds),
			InitialDelaySeconds: int32(data.ReadinessProbe.InitialDelaySeconds),
		}
	}
	if data.StartupProbe != nil && data.StartupProbe.Path != "" && data.StartupProbe.Port > 0 {
		param.StartupProbe = &pbSpec.HttpGetProbe{
			Path:                data.StartupProbe.Path,
			Port:                int32(data.StartupProbe.Port),
			FailureThreshold:    int32(data.StartupProbe.FailureThreshold),
			PeriodSeconds:       int32(data.StartupProbe.PeriodSeconds),
			InitialDelaySeconds: int32(data.StartupProbe.InitialDelaySeconds),
		}
	}
	_, err := r.cli.UpsertDeploymentConfig(ctx, param)
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

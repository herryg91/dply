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
func (r *repository) GetEnvar(env, name string) (*entity.Envar, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetEnvar(ctx, &pbSpec.GetEnvarReq{Env: env, Name: name})
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

	envar := &entity.Envar{Env: env, Name: name, Variables: map[string]interface{}{}}
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

func (r *repository) GetScale(env, name string) (*entity.Scale, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetScale(ctx, &pbSpec.GetScaleReq{Env: env, Name: name})
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

func (r *repository) GetPort(env, name string) (*entity.Port, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetPort(ctx, &pbSpec.GetPortReq{Env: env, Name: name})
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
			Name:     p.PortName,
			Port:     int(p.Port),
			Protocol: entity.PortType(p.Protocol),
		})
	}
	port := &entity.Port{Env: env, Name: name, Ports: portSpecs}
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
			PortName: p.Name,
			Port:     int32(p.Port),
			Protocol: string(p.Protocol),
		})
	}
	_, err := r.cli.UpsertPort(ctx, &pbSpec.UpsertPortReq{
		Env:   data.Env,
		Name:  data.Name,
		Ports: portParam,
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

func (r *repository) GetAffinity(env, name string) (*entity.Affinity, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	resp, err := r.cli.GetAffinity(ctx, &pbSpec.GetAffinityReq{Env: env, Name: name})
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

	for _, a := range resp.NodeAffinity {
		nodeAffinity = append(nodeAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAffinity {
		podAffinity = append(podAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range resp.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, entity.AffinityTerm{Mode: entity.AffinityMode(a.Mode), Key: a.Key, Operator: entity.AffinityOperator(a.Operator), Values: a.Values, Weight: int(a.Weight), TopologyKey: a.TopologyKey})
	}

	affinity := &entity.Affinity{Env: env, Name: name, NodeAffinity: nodeAffinity, PodAffinity: podAffinity, PodAntiAffinity: podAntiAffinity}

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

	for _, a := range data.NodeAffinity {
		nodeAffinity = append(nodeAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}

	for _, a := range data.PodAffinity {
		podAffinity = append(podAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}
	for _, a := range data.PodAntiAffinity {
		podAntiAffinity = append(podAntiAffinity, &pbSpec.AffinityTerm{Mode: string(a.Mode), Key: a.Key, Operator: string(a.Operator), Values: a.Values, Weight: int32(a.Weight), TopologyKey: a.TopologyKey})
	}

	_, err := r.cli.UpsertAffinity(ctx, &pbSpec.UpsertAffinityReq{
		Env:             data.Env,
		Name:            data.Name,
		NodeAffinity:    nodeAffinity,
		PodAffinity:     podAffinity,
		PodAntiAffinity: podAntiAffinity,
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

package deployment_config_usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	deployment_config_repo repository.DeploymentConfigRepository
}

func New(deployment_config_repo repository.DeploymentConfigRepository) UseCase {
	return &usecase{deployment_config_repo: deployment_config_repo}
}

func (uc *usecase) Get(project, env, name string) (*entity.DeploymentConfig, error) {
	resp, err := uc.deployment_config_repo.Get(project, env, name)
	if err != nil {
		if errors.Is(err, repository.ErrDeploymentConfigNotFound) {
			return &entity.DeploymentConfig{
				Project:        project,
				Env:            env,
				Name:           name,
				LivenessProbe:  nil,
				ReadinessProbe: nil,
				StartupProbe:   nil,
			}, nil
		} else {
			return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	return resp, nil
}

func (uc *usecase) Upsert(data entity.DeploymentConfig) error {
	data.Env = strings.ToLower(data.Env)
	data.Name = strings.ToLower(data.Name)

	err := uc.deployment_config_repo.Upsert(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return nil
}

package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrDeploymentConfigNotFound = errors.New("Deployment Config spec not found")

type DeploymentConfigRepository interface {
	Get(project, env, name string) (*entity.DeploymentConfig, error)
	Upsert(data entity.DeploymentConfig) error
}

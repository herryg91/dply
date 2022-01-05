package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrDeploymentNotFound = errors.New("Current deployment not found")

type DeploymentRepository interface {
	Get(project, env, name string) (*entity.Deployment, error)
	Create(in entity.Deployment) error
}

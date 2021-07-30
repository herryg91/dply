package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrK8sDeploymentNotFound = errors.New("deployment.apps not found")
var ErrK8sServiceNotFound = errors.New("service not found")
var ErrK8sHPANotFound = errors.New("hpa not found")

type K8sRepository interface {
	CheckDeploymentExist(env, name string) error
	CreateDeployment(param entity.Deployment) error
	UpdateDeployment(param entity.Deployment) error

	CheckServiceExist(env, name string) error
	CreateService(param entity.Deployment) error
	UpdateService(param entity.Deployment) error

	CheckHPAExist(env, name string) error
	CreateHPA(param entity.Deployment) error
	UpdateHPA(param entity.Deployment) error
}

package deploy_usecase

import (
	"errors"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	DeployImage(project string, env string, name string, digest string, createdBy int) error
	Redeploy(project string, env string, serviceName string, createdBy int) error
}

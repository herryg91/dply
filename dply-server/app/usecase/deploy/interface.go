package deploy_usecase

import (
	"errors"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	DeployImage(env string, name string, digest string, createdBy int) error
	Redeploy(env string, serviceName string, createdBy int) error
}

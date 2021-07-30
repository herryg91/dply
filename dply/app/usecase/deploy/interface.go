package deploy_usecase

import (
	"errors"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("Unauthorized action")

type UseCase interface {
	Deploy(env, name, digest string) error
	Redeploy(env, name string) error
}

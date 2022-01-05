package project_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrProjectAlreadyExist = errors.New("Project is already exist")
var ErrProjectNameInvalidFormat = errors.New("Parameter name is invalid format")

type UseCase interface {
	Get() ([]entity.Project, error)
	Create(name, description string) error
	Delete(name string) error
}

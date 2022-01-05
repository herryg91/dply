package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrProjectDuplicate = errors.New("Duplicate project name")

type ProjectRepository interface {
	GetAll() ([]entity.Project, error)
	Create(p entity.Project) error
	DeleteByName(name string) error
}

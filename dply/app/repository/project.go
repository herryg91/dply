package repository

import (
	"github.com/herryg91/dply/dply/entity"
)

type ProjectRepository interface {
	GetAll() ([]entity.Project, error)
	Create(p entity.Project) error
	Delete(name string) error
}

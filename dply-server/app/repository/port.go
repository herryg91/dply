package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrPortNotFound = errors.New("Port spec not found")
var ErrPortTemplateNotFound = errors.New("Port template not found")

type PortRepository interface {
	Get(env, name string) (*entity.Port, error)
	Upsert(data entity.Port) error
	GetPortByTemplate(templateName string) (*entity.PortTemplate, error)
	UpsertPortByTemplate(data entity.PortTemplate) error
}

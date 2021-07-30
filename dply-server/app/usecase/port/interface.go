package port_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	Get(env string, name string) (*entity.Port, error)
	Upsert(data entity.Port) error
	GetTemplate(templateName string) (*entity.PortTemplate, error)
	UpsertTemplate(data entity.PortTemplate) error
}

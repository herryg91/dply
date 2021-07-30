package affinity_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	Get(env string, name string) (*entity.Affinity, error)
	Upsert(data entity.Affinity) error
	GetTemplate(templateName string) (*entity.AffinityTemplate, error)
	UpsertTemplate(data entity.AffinityTemplate) error
}

package scale_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	Get(env string, name string) (*entity.Scale, error)
	Upsert(data entity.Scale) error
}

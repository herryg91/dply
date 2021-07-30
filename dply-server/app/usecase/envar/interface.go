package envar_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")

type UseCase interface {
	Get(env string, name string) (*entity.Envar, error)
	Upsert(data entity.Envar) error
}

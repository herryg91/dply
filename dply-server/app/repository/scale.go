package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrScaleNotFound = errors.New("Service's scaling config not found")

type ScaleRepository interface {
	Get(project, env, name string) (*entity.Scale, error)
	Upsert(data entity.Scale) error
}

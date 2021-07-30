package repository

import (
	"github.com/herryg91/dply/dply/entity"
)

type SpecRepository interface {
	GetEnvar(env, name string) (*entity.Envar, error)
	UpsertEnvar(data entity.Envar) error

	GetScale(env, name string) (*entity.Scale, error)
	UpsertScale(data entity.Scale) error

	GetPort(env, name string) (*entity.Port, error)
	UpsertPort(data entity.Port) error

	GetAffinity(env, name string) (*entity.Affinity, error)
	UpsertAffinity(data entity.Affinity) error
}

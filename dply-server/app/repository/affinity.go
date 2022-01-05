package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrAffinityNotFound = errors.New("Affinity spec not found")
var ErrAffinityTemplateNotFound = errors.New("Affinity template not found")

type AffinityRepository interface {
	Get(project, env, name string) (*entity.Affinity, error)
	Upsert(data entity.Affinity) error
	GetAffinityByTemplate(templateName string) (*entity.AffinityTemplate, error)
	UpsertAffinityByTemplate(data entity.AffinityTemplate) error
}

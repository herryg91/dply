package deployment_config_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("Unauthorized action")

type UseCase interface {
	Get(project, env, name string) (*entity.DeploymentConfig, error)
	Upsert(data entity.DeploymentConfig) error
	UpsertViaEditor(project, env, name string, editorApp editor.EditorApp) (bool, error)
}

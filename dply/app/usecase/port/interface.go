package port_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("Unauthorized action")

type UseCase interface {
	Get(env, name string) (*entity.Port, error)
	Upsert(data entity.Port) error
	UpsertViaEditor(env, name string, editorApp editor.EditorApp) (bool, error)
}

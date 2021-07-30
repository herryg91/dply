package envar_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("Unauthorized action")

type UseCase interface {
	Get(env, name string) (*entity.Envar, error)
	Upsert(data entity.Envar) error
	UpsertViaEditor(env, name string, editorApp editor.EditorApp) (bool, error)
}

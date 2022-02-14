package image_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("Unauthorized action")

type UseCase interface {
	Create(project, name, tag_prefix, description string, build_args map[string]*string) error
	Add(project, repoName, image, description string) error
	Remove(repoName, digest string) error
	GetList(project, repoName string, page, size int) ([]entity.ContainerImage, error)
}

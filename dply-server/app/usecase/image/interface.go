package image_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrImageNotFound = errors.New("Image not found")

type UseCase interface {
	Get(repository string, page, size int) ([]entity.Image, error)
	Add(repository, fullImage, description string, createdBy int) error
	Remove(digest string) error
	GetByDigest(digest string) (*entity.Image, error)
}

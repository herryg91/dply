package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrImageDigestDuplicate = errors.New("image digest is duplicate")
var ErrImageNotFound = errors.New("image not found")

type ImageRepository interface {
	Search(repoName string, limit, offset int, createdAtDesc bool) ([]entity.Image, error)
	Create(in entity.Image) error
	Delete(digest string) error
	GetByDigest(digest string) (*entity.Image, error)
	GetByImage(image string) (*entity.Image, error)
}

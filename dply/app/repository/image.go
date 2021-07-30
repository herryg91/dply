package repository

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
)

var ErrUnauthorizedAdmin = errors.New("unauthorized action (require admin access)")

type ImageRepository interface {
	Add(repoName, image, description string) error
	Remove(repoName, digest string) error
	Get(repoName string, page, size int) ([]entity.ContainerImage, error)
}

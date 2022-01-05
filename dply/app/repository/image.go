package repository

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
)

var ErrUnauthorizedAdmin = errors.New("unauthorized action (require admin access)")

type ImageRepository interface {
	Add(project, repoName, image, description string) error
	Remove(repoName, digest string) error
	Get(project, repoName string, page, size int) ([]entity.ContainerImage, error)

	// Build Docker Image
	BuildImage(repo_full_name string, src string) (docker_image_ids []string, err error)
	// Push Image to Registry
	PushImage(image_tag_name string) (digest string, err error)
	DeleteImage(image_id string) error
}

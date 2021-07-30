package image_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply/app/repository"
	"github.com/herryg91/dply/dply/entity"
)

type usecase struct {
	repo repository.ImageRepository
}

func New(repo repository.ImageRepository) UseCase {
	return &usecase{repo: repo}
}
func (uc *usecase) Add(repoName, image, description string) error {
	err := uc.repo.Add(repoName, image, description)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}
func (uc *usecase) Remove(repoName, digest string) error {
	err := uc.repo.Remove(repoName, digest)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		} else if errors.Is(err, repository.ErrUnauthorizedAdmin) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login (require admin access)")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}
func (uc *usecase) GetList(repoName string, page, size int) ([]entity.ContainerImage, error) {
	resp, err := uc.repo.Get(repoName, page, size)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return []entity.ContainerImage{}, fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return []entity.ContainerImage{}, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

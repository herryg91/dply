package image_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	repo           repository.ImageRepository
	deplyment_repo repository.DeploymentRepository
}

func New(repo repository.ImageRepository) UseCase {
	return &usecase{repo: repo}
}

func (uc *usecase) Get(project, repositoryName string, page, size int) ([]entity.Image, error) {
	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}

	limit := size
	offset := (page - 1) * size

	resp, err := uc.repo.Search(project, repositoryName, limit, offset, true)
	if err != nil {
		return []entity.Image{}, fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	// Create Notes (last deployment environment)
	deployments, err := uc.deplyment_repo.GetLatestDeploymentGroupByName(project, repositoryName)
	if err != nil {
		return []entity.Image{}, fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	map_images_env := map[string]string{}
	for _, d := range deployments {
		if _, ok := map_images_env[d.DeploymentImage.Digest]; !ok {
			map_images_env[d.DeploymentImage.Digest] = d.Env
		} else {
			map_images_env[d.DeploymentImage.Digest] += " | " + d.Env
		}
	}

	for i, r := range resp {
		if notes, ok := map_images_env[r.Digest]; ok {
			resp[i].Notes = notes
		}
	}

	return resp, nil
}

var ErrInvalidImageFormat = errors.New("invalid image format. valid format: <repo_name>@<digest>. Example: gcr.io/repo@sha256:xxxx")
var ErrImageAlreadyExist = errors.New("image is already exist")
var ErrDigestAlreadyExist = errors.New("digest is already exist, potential race condition. please retry")

func (uc *usecase) Add(project, repositoryName, fullImage, description string, createdBy int) error {
	fullDigest, err := uc.imageToDigest(fullImage)
	if err != nil {
		return ErrInvalidImageFormat
	}

	shortDigest, err := uc.generateShortDigest(fullDigest)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	// Check if digest can be write to db
	_, err = uc.repo.GetByImage(fullImage)
	if err != nil {
		if !errors.Is(err, repository.ErrImageNotFound) {
			return fmt.Errorf("%w: %v", ErrUnexpected, err)
		}
	} else {
		return ErrImageAlreadyExist
	}

	// Write to db
	err = uc.repo.Create(entity.Image{
		Project:     project,
		Digest:      shortDigest,
		Image:       fullImage,
		Repository:  repositoryName,
		Description: description,
		CreatedBy:   createdBy,
	})
	if err != nil {
		if errors.Is(err, repository.ErrImageDigestDuplicate) {
			return ErrDigestAlreadyExist
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return nil
}

func (uc *usecase) Remove(digest string) error {
	err := uc.repo.Delete(digest)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return nil
}

func (uc *usecase) GetByDigest(digest string) (*entity.Image, error) {
	resp, err := uc.repo.GetByDigest(digest)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return resp, nil
}

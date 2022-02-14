package image_usecase

import (
	"errors"
	"fmt"
	"os"

	"github.com/herryg91/dply/dply/app/repository"
	"github.com/herryg91/dply/dply/entity"
)

type usecase struct {
	repo repository.ImageRepository
}

func New(repo repository.ImageRepository) UseCase {
	return &usecase{repo: repo}
}
func (uc *usecase) Add(project, repoName, image, description string) error {
	err := uc.repo.Add(project, repoName, image, description)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

func (uc *usecase) Create(project, name, tag_prefix, description string, build_args map[string]*string) error {
	// Preparation
	current_folder, _ := os.Getwd()
	if _, err := os.Stat(current_folder + "/Dockerfile"); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%w: %v", ErrUnexpected, "Dockerfile not exist")
	}

	if len(tag_prefix) > 0 && string(tag_prefix[len(tag_prefix)-1]) != "/" {
		tag_prefix += "/"
	}
	repo_full_name := tag_prefix + name

	fmt.Println("---------- 1. Building Docker Image ----------")
	docker_image_ids, err := uc.repo.BuildImage(repo_full_name, current_folder, build_args)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	fmt.Println("---------- 2. Push Image to Registry Server ----------")
	digest, err := uc.repo.PushImage(repo_full_name)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	fmt.Println("Successfully push image, digest: ", digest)

	fmt.Println("---------- 3. Add Image to Dply ----------")
	err = uc.Add(project, name, repo_full_name+"@"+digest, description)
	if err != nil {
		return err
	}
	fmt.Println("Push Image done")

	fmt.Println("---------- 4. Clean Up Docker Images ----------")
	for _, img := range docker_image_ids {
		fmt.Println("Delete " + img + "")
		err = uc.repo.DeleteImage(img)
		if err != nil {
			if errors.Is(err, repository.ErrUnauthorized) {
				return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
			}
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}
	fmt.Println()
	fmt.Println("Done")
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
func (uc *usecase) GetList(project, repoName string, page, size int) ([]entity.ContainerImage, error) {
	resp, err := uc.repo.Get(project, repoName, page, size)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return []entity.ContainerImage{}, fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return []entity.ContainerImage{}, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

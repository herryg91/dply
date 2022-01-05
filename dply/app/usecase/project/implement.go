package project_usecase

import (
	"fmt"

	"github.com/herryg91/dply/dply/app/repository"
	"github.com/herryg91/dply/dply/entity"
)

type usecase struct {
	repo repository.ProjectRepository
}

func New(repo repository.ProjectRepository) UseCase {
	return &usecase{repo: repo}
}

func (uc *usecase) Get() ([]entity.Project, error) {
	resp, err := uc.repo.GetAll()
	if err != nil {
		return []entity.Project{}, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) Create(p entity.Project) error {
	err := uc.repo.Create(p)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

func (uc *usecase) Delete(name string) error {
	err := uc.repo.Delete(name)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

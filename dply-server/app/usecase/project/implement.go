package project_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	project_repo repository.ProjectRepository
}

func New(project_repo repository.ProjectRepository) UseCase {
	return &usecase{project_repo: project_repo}
}

func (uc *usecase) Get() ([]entity.Project, error) {
	res, err := uc.project_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return res, err
}

func (uc *usecase) Create(name, description string) error {
	p := entity.Project{Name: name, Description: description}
	err := p.ValidateName()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrProjectNameInvalidFormat, err.Error())
	}
	err = uc.project_repo.Create(p)
	if err != nil {
		if errors.Is(err, repository.ErrProjectDuplicate) {
			return ErrProjectAlreadyExist
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

func (uc *usecase) Delete(name string) error {
	err := uc.project_repo.DeleteByName(name)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

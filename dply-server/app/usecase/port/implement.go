package port_usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	port_repo repository.PortRepository
}

func New(port_repo repository.PortRepository) UseCase {
	return &usecase{port_repo: port_repo}
}

func (uc *usecase) Get(project, env, name string) (*entity.Port, error) {
	resp, err := uc.port_repo.Get(project, env, name)
	useDefault := false
	if err != nil {
		if errors.Is(err, repository.ErrPortNotFound) {
			useDefault = true
		} else {
			return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	if useDefault {
		defaultTmpl, err := uc.GetTemplate("default")
		if err != nil {
			return nil, err
		}

		return defaultTmpl.ToPortEntity(env, name), nil
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.Port) error {
	data.Env = strings.ToLower(data.Env)
	data.Name = strings.ToLower(data.Name)

	err := uc.port_repo.Upsert(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return nil
}

func (uc *usecase) GetTemplate(templateName string) (*entity.PortTemplate, error) {
	resp, err := uc.port_repo.GetPortByTemplate(templateName)
	if err != nil {
		if errors.Is(err, repository.ErrPortTemplateNotFound) {
			return entity.PortTemplate{}.DefaultPortTemplate(), nil
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) UpsertTemplate(data entity.PortTemplate) error {
	err := uc.port_repo.UpsertPortByTemplate(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

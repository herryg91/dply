package affinity_usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	affinity_repo repository.AffinityRepository
}

func New(affinity_repo repository.AffinityRepository) UseCase {
	return &usecase{affinity_repo: affinity_repo}
}

func (uc *usecase) Get(env string, name string) (*entity.Affinity, error) {
	resp, err := uc.affinity_repo.Get(env, name)
	useDefault := false
	if err != nil {
		if errors.Is(err, repository.ErrAffinityNotFound) {
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

		return defaultTmpl.ToAffinityEntity(env, name), nil
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.Affinity) error {
	data.Env = strings.ToLower(data.Env)
	data.Name = strings.ToLower(data.Name)

	err := uc.affinity_repo.Upsert(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return nil
}

func (uc *usecase) GetTemplate(templateName string) (*entity.AffinityTemplate, error) {
	resp, err := uc.affinity_repo.GetAffinityByTemplate(templateName)
	if err != nil {
		if errors.Is(err, repository.ErrAffinityTemplateNotFound) {
			return entity.AffinityTemplate{}.DefaultAffinityTemplate(), nil
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) UpsertTemplate(data entity.AffinityTemplate) error {
	err := uc.affinity_repo.UpsertAffinityByTemplate(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

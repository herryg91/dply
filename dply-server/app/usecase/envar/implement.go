package envar_usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	envar_repo repository.EnvarRepository
}

func New(envar_repo repository.EnvarRepository) UseCase {
	return &usecase{envar_repo: envar_repo}
}

func (uc *usecase) Get(env string, name string) (*entity.Envar, error) {
	resp, err := uc.envar_repo.Get(env, name)
	if err != nil {
		if errors.Is(err, repository.ErrEnvarNotFound) {
			return entity.Envar{}.DefaultEnvar(env, name), nil
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.Envar) error {
	data.Env = strings.ToLower(data.Env)
	data.Name = strings.ToLower(data.Name)

	err := uc.envar_repo.Upsert(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return nil
}

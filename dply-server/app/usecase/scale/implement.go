package scale_usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	scale_repo repository.ScaleRepository
}

func New(scale_repo repository.ScaleRepository) UseCase {
	return &usecase{scale_repo: scale_repo}
}

func (uc *usecase) Get(env string, name string) (*entity.Scale, error) {
	resp, err := uc.scale_repo.Get(env, name)
	if err != nil {
		if errors.Is(err, repository.ErrScaleNotFound) {
			return entity.Scale{}.DefaultScale(env, name), nil
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return resp, nil
}

func (uc *usecase) Upsert(data entity.Scale) error {
	data.Env = strings.ToLower(data.Env)
	data.Name = strings.ToLower(data.Name)

	err := uc.scale_repo.Upsert(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return nil
}

package deploy_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply/app/repository"
)

type usecase struct {
	repo repository.DeployRepository
}

func New(repo repository.DeployRepository) UseCase {
	return &usecase{repo: repo}
}
func (uc *usecase) Deploy(project, env, name, digest string) error {
	err := uc.repo.Deploy(project, env, name, digest)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}
func (uc *usecase) Redeploy(project, env, name string) error {
	err := uc.repo.Redeploy(project, env, name)
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return nil
}

package server_usecase

import (
	"github.com/herryg91/dply/dply/app/repository"
)

type usecase struct {
	repo repository.ServerRepository
}

func New(repo repository.ServerRepository) UseCase {
	return &usecase{repo: repo}
}

func (uc *usecase) Status() bool {
	return uc.repo.Status()
}

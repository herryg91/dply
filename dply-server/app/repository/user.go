package repository

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserDuplicate = errors.New("user duplicate")
var ErrUserInactive = errors.New("user is inactive")

type UserRepository interface {
	GetById(id int) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByToken(token string) (*entity.User, error)
	Create(req entity.User) error
	Edit(req entity.User) error
	EditStatus(email string, status bool) error
	EditPassword(email string, password string, newToken string) error
}

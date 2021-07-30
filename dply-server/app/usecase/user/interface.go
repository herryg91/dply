package user_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply-server/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUserNotFound = errors.New("User not found")

type UseCase interface {
	Login(email, password string) (*entity.User, error)
	Register(email, password string, userType entity.UserType, name string) error
	Edit(email, password string, userType entity.UserType, name string) error
	GetByToken(token string) (*entity.User, error)
	EditStatusToActive(email string) error
	EditStatusToInactive(email string) error
	EditPassword(email, oldPassword, newPassword string) error
	IsAdmin(token string) bool
}

package repository

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
)

var ErrUserUnexpected = errors.New("unexpected error")
var ErrUserInvalidPassword = errors.New("invalid password")
var ErrUserNotRegistered = errors.New("email isn't registered")
var ErrUserDuplicate = errors.New("duplicate user")
var ErrUserInactive = errors.New("user is inactive")
var ErrUserUnauthorized = errors.New("unauthorized action")

type UserRepository interface {
	Login(email, password string) (*entity.User, error)
	GetCurrentLogin() (*entity.User, error)
	CheckLogin() error
}

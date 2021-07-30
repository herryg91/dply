package auth_usecase

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnauthorized = errors.New("unauthorized action")

type UseCase interface {
	Login(email, password string) error
	Logout()
	GetStatus() (isLogin bool, userData *entity.User)
	CheckLogin() error
}

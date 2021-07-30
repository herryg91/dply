package user_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
	"github.com/herryg91/dply/dply-server/pkg/helpers"
	"github.com/herryg91/dply/dply-server/pkg/password"
)

type usecase struct {
	user_repo    repository.UserRepository
	password_svc password.Password
}

func New(user_repo repository.UserRepository, password_svc password.Password) UseCase {
	return &usecase{user_repo: user_repo, password_svc: password_svc}
}

var ErrUserInvalidPassword = errors.New("Invalid password")
var ErrUserInactive = errors.New("User is inactive")

func (uc *usecase) Login(email, password string) (*entity.User, error) {
	u, err := uc.user_repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		} else if errors.Is(err, repository.ErrUserInactive) {
			return nil, ErrUserInactive
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	if !uc.password_svc.Check(password, u.Password) {
		return nil, ErrUserInvalidPassword
	}

	return u, nil
}

var ErrUserAlreadyExist = errors.New("User already exist")

func (uc *usecase) Register(email, password string, userType entity.UserType, name string) error {
	hashedPassword, err := uc.password_svc.Hash(password)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	err = uc.user_repo.Create(entity.User{
		Email:    email,
		Password: hashedPassword,
		UserType: userType,
		Name:     name,
		Token:    helpers.RandomString(64),
	})
	if err != nil {
		if errors.Is(err, repository.ErrUserDuplicate) {
			return fmt.Errorf("%w: %v", ErrUserAlreadyExist, err)
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	return nil
}

func (uc *usecase) Edit(email, password string, userType entity.UserType, name string) error {
	u, err := uc.user_repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	if password != "" {
		hashedPassword, err := uc.password_svc.Hash(password)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err)
		}

		u.Password = hashedPassword
		u.Token = helpers.RandomString(64)
	}

	if userType == entity.UserType_Admin || userType == entity.UserType_User {
		u.UserType = userType
	}

	if name != "" {
		u.Name = name
	}

	err = uc.user_repo.Edit(*u)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return nil
}
func (uc *usecase) GetByToken(token string) (*entity.User, error) {
	u, err := uc.user_repo.GetByToken(token)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	return u, nil
}

func (uc *usecase) EditPassword(email, oldPassword, newPassword string) error {
	u, err := uc.user_repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		} else if errors.Is(err, repository.ErrUserInactive) {
			return ErrUserInactive
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	if !uc.password_svc.Check(oldPassword, u.Password) {
		return ErrUserInvalidPassword
	}

	hashedNewPassword, err := uc.password_svc.Hash(newPassword)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	err = uc.user_repo.EditPassword(email, hashedNewPassword, helpers.RandomString(64))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return nil
}

func (uc *usecase) EditStatusToActive(email string) error {
	err := uc.user_repo.EditStatus(email, true)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	return nil
}
func (uc *usecase) EditStatusToInactive(email string) error {
	if email == "admin@dply.com" {
		return fmt.Errorf("%v", "Cannot set to inactive root admin: admin@dply.com")
	}
	err := uc.user_repo.EditStatus(email, false)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	return nil
}

func (uc *usecase) IsAdmin(token string) bool {
	u, err := uc.user_repo.GetByToken(token)
	if err != nil {
		return false
	}

	return u.UserType == entity.UserType_Admin
}

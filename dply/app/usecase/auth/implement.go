package auth_usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/herryg91/dply/dply/app/repository"
	"github.com/herryg91/dply/dply/entity"
)

type usecase struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UseCase {
	return &usecase{repo: repo}
}

var ErrLoginFailed = errors.New("Failed to login")

func (uc *usecase) Login(email, password string) error {
	u, err := uc.repo.Login(email, password)
	if err != nil {
		if errors.Is(err, repository.ErrUserInvalidPassword) {
			return fmt.Errorf("%w: %v", ErrLoginFailed, err)
		} else if errors.Is(err, repository.ErrUserNotRegistered) {
			return fmt.Errorf("%w: %v", ErrLoginFailed, err)
		} else if errors.Is(err, repository.ErrUserInactive) {
			return fmt.Errorf("%v", "User is inactive")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	if _, err := os.Stat(dplyDir); os.IsNotExist(err) {
		os.Mkdir(dplyDir, 0755)
	}

	userJson, _ := json.Marshal(&u)
	err = ioutil.WriteFile(dplyDir+"/user.json", userJson, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}

	return nil
}

func (uc *usecase) Logout() {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	userFileLoc := dplyDir + "/user.json"
	if _, err := os.Stat(userFileLoc); !os.IsNotExist(err) {
		os.Remove(userFileLoc)
	}

	return
}

func (uc *usecase) GetStatus() (isLogin bool, userData *entity.User) {
	isLogin, userData = false, nil

	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	userFileLoc := dplyDir + "/user.json"
	if _, err := os.Stat(userFileLoc); os.IsNotExist(err) {
		return
	}

	userDataFile := &entity.User{}
	userFileData, _ := ioutil.ReadFile(userFileLoc)
	_ = json.Unmarshal(userFileData, &userDataFile)
	if userDataFile != nil {
		if userDataFile.Token != "" {
			isLogin = true
		}
	}

	var err error
	userData, err = uc.repo.GetCurrentLogin()
	if err != nil {
		return false, nil
	}
	return
}

func (uc *usecase) CheckLogin() error {
	err := uc.repo.CheckLogin()
	if err != nil {
		if errors.Is(err, repository.ErrUnauthorized) {
			return fmt.Errorf("%w: %v", ErrUnauthorized, "You are not login")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return nil
}

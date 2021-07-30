package user_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbUser "github.com/herryg91/dply/dply/clients/grst/user"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type repository struct {
	cli pbUser.UserApiClient
}

func New(cli pbUser.UserApiClient) repository_intf.UserRepository {
	return &repository{cli}
}

func (r *repository) Login(email, password string) (*entity.User, error) {
	u, err := r.cli.Login(context.Background(), &pbUser.LoginReq{
		Email:    email,
		Password: password,
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}

		if grsterr.Code == 10002 {
			return nil, repository_intf.ErrUserInvalidPassword
		} else if grsterr.Code == 10001 {
			return nil, repository_intf.ErrUserNotRegistered
		} else if grsterr.Code == 10003 {
			return nil, repository_intf.ErrUserInactive
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return &entity.User{
		Name:     u.Name,
		Usertype: u.Usertype,
		Email:    u.Email,
		Token:    u.Token,
	}, nil
}

func (r *repository) GetCurrentLogin() (*entity.User, error) {
	uctx := entity.User{}.FromFile()
	if uctx == nil {
		return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": uctx.Token}))

	u, err := r.cli.GetCurrentLogin(ctx, &emptypb.Empty{})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return nil, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return nil, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return nil, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return &entity.User{
		Name:     u.Name,
		Usertype: u.Usertype,
		Email:    u.Email,
		Token:    u.Token,
	}, nil
}

func (r *repository) CheckLogin() error {
	uctx := entity.User{}.FromFile()
	if uctx == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": uctx.Token}))

	_, err := r.cli.GetCurrentLogin(ctx, &emptypb.Empty{})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}

		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

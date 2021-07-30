package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	pbUser "github.com/herryg91/dply/dply-server/handler/grst/user"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"google.golang.org/grpc/codes"
)

type handlerUser struct {
	user_uc user_usecase.UseCase
	pbUser.UnimplementedUserApiServer
}

func NewUserHandler(user_uc user_usecase.UseCase) pbUser.UserApiServer {
	return &handlerUser{user_uc: user_uc}
}

func (h *handlerUser) Login(ctx context.Context, req *pbUser.LoginReq) (*pbUser.User, error) {
	if err := pbUser.ValidateRequest(req); err != nil {
		return nil, err
	}
	u, err := h.user_uc.Login(req.Email, req.Password)
	if err != nil {
		if err != nil {
			if errors.Is(err, user_usecase.ErrUserNotFound) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 10001, err.Error(), &grst_errors.ErrorDetail{Code: 1, Field: "email", Message: "email `" + req.Email + "` isn't registered"})
			} else if errors.Is(err, user_usecase.ErrUserInvalidPassword) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 10002, err.Error(), &grst_errors.ErrorDetail{Code: 1, Field: "password", Message: "invalid password"})
			} else if errors.Is(err, user_usecase.ErrUserInactive) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 10003, err.Error())
			}
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10000, err.Error())
		}
	}
	return &pbUser.User{
		Name:     u.Name,
		Usertype: string(u.UserType),
		Email:    u.Email,
		Token:    u.Token,
	}, nil
}

func (h *handlerUser) GetCurrentLogin(ctx context.Context, req *empty.Empty) (*pbUser.User, error) {
	uctx := interceptor.ExtractMustLoginContext(ctx)
	if uctx == nil || uctx.Token == "" {
		return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 13004, "You are not login")
	}

	u, err := h.user_uc.GetByToken(uctx.Token)
	if err != nil {
		if err != nil {
			if errors.Is(err, user_usecase.ErrUserNotFound) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 13001, err.Error())
			} else if errors.Is(err, user_usecase.ErrUserInactive) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 13003, err.Error())
			}
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 13000, err.Error())
		}
	}

	return &pbUser.User{
		Name:     u.Name,
		Usertype: string(u.UserType),
		Email:    u.Email,
		Token:    u.Token,
	}, nil
}

func (h *handlerUser) UpdatePassword(ctx context.Context, req *pbUser.UpdatePasswordReq) (*empty.Empty, error) {
	if err := pbUser.ValidateRequest(req); err != nil {
		return nil, err
	}

	uctx := interceptor.ExtractMustLoginContext(ctx)
	if uctx == nil || uctx.Token == "" {
		return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 14001, "You are not login")
	} else if uctx.Email != req.Email {
		return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 14002, "You cannot change other user's password. Please login as "+req.Email+" to change the password")
	}

	err := h.user_uc.EditPassword(req.Email, req.OldPassword, req.NewPassword)
	if err != nil {
		if err != nil {
			if errors.Is(err, user_usecase.ErrUserNotFound) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 14003, err.Error(), &grst_errors.ErrorDetail{Code: 1, Field: "email", Message: "email `" + req.Email + "` isn't registered"})
			} else if errors.Is(err, user_usecase.ErrUserInvalidPassword) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 14004, err.Error(), &grst_errors.ErrorDetail{Code: 1, Field: "old_password", Message: "invalid old password"})
			} else if errors.Is(err, user_usecase.ErrUserInactive) {
				return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 14005, err.Error())
			}
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 14000, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

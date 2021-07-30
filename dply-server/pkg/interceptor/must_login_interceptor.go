package interceptor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	"github.com/herryg91/dply/dply-server/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func MustLoginInterceptor(user_uc user_usecase.UseCase, fullMethods []string) grpc.UnaryServerInterceptor {
	fullMethodsMap := map[string]bool{}
	for _, fm := range fullMethods {
		fullMethodsMap[fm] = true
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if _, ok := fullMethodsMap[info.FullMethod]; !ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 98001, "Failed to parse context")
		} else if len(md["authorization"]) <= 0 {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 98002, "Authorization on header is required")
		}

		token := md["authorization"][0]

		u, err := user_uc.GetByToken(token)
		if err != nil {
			if errors.Is(err, user_usecase.ErrUserNotFound) {
				return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 98003, "You are not login")
			}
			return nil, err
		}

		ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
			"user.id":       strconv.Itoa(u.Id),
			"user.name":     u.Name,
			"user.email":    u.Email,
			"user.usertype": string(u.UserType),
			"user.token":    u.Token,
		}))
		return handler(ctx, req)
	}
}

func ExtractMustLoginContext(ctx context.Context) *entity.User {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	resp := &entity.User{}
	if val := md.Get("user.id"); len(val) > 0 {
		fmt.Sscanf(val[0], "%d", &resp.Id)
	}
	if val := md.Get("user.name"); len(val) > 0 {
		resp.Name = val[0]
	}
	if val := md.Get("user.email"); len(val) > 0 {
		resp.Email = val[0]
	}
	if val := md.Get("user.usertype"); len(val) > 0 {
		resp.UserType = entity.UserType(val[0])
	}
	if val := md.Get("user.token"); len(val) > 0 {
		resp.Token = val[0]
	}
	return resp
}

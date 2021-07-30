package interceptor

import (
	"context"
	"net/http"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func LoginAsAdminInterceptor(user_uc user_usecase.UseCase, fullMethods []string) grpc.UnaryServerInterceptor {
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

		if !user_uc.IsAdmin(token) {
			return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 98003, "Action is prohibited (require admin access)")
		}

		return handler(ctx, req)
	}
}

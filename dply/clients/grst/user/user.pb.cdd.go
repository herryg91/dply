// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: user.proto
package user

import (
	"net/http"
	"strings"

	"github.com/herryg91/cdd/grst"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc"

	"github.com/mcuadros/go-defaults"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"gopkg.in/validator.v2"
)

type fullMethods struct {
	UserApi_Login           string
	UserApi_GetCurrentLogin string
	UserApi_UpdatePassword  string
}

var FullMethods = fullMethods{
	UserApi_Login:           "/user.UserApi/Login",
	UserApi_GetCurrentLogin: "/user.UserApi/GetCurrentLogin",
	UserApi_UpdatePassword:  "/user.UserApi/UpdatePassword",
}

var NeedAuthFullMethods = []string{}

type AuthConfig struct {
	NeedAuth bool
	Roles    []string
}

var AuthConfigFullMethods = map[string]AuthConfig{
	"/user.UserApi/Login":           AuthConfig{NeedAuth: false, Roles: []string{"*"}},
	"/user.UserApi/GetCurrentLogin": AuthConfig{NeedAuth: false, Roles: []string{"*"}},
	"/user.UserApi/UpdatePassword":  AuthConfig{NeedAuth: false, Roles: []string{"*"}},
}

var NeedApiKeyFullMethods = []string{}

func ValidateRequest(req interface{}) error {
	defaults.SetDefaults(req)
	if errs := validator.Validate(req); errs != nil {
		validateError := []*grst_errors.ErrorDetail{}
		for field, err := range errs.(validator.ErrorMap) {
			errMessage := strings.Replace(err.Error(), "{field}", field, -1)
			validateError = append(validateError, &grst_errors.ErrorDetail{Code: 999, Field: field, Message: errMessage})
		}
		return grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 999, "Validation Error", validateError...)
	}

	return nil
}

/*==================== UserApi Section ====================*/

func RegisterUserApiGrstServer(grpcRestServer *grst.Server, hndl UserApiServer) {

	forward_UserApi_Login_0 = grpcRestServer.GetForwardResponseMessage()

	forward_UserApi_GetCurrentLogin_0 = grpcRestServer.GetForwardResponseMessage()

	forward_UserApi_UpdatePassword_0 = grpcRestServer.GetForwardResponseMessage()

	RegisterUserApiServer(grpcRestServer.GetGrpcServer(), hndl)
	grpcRestServer.RegisterRestHandler(RegisterUserApiHandler)
}

func NewUserApiGrstClient(serverHost string, creds *credentials.TransportCredentials, dialOpts ...grpc.DialOption) (UserApiClient, error) {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)))
	opts = append(opts, grpc.WithMaxMsgSize(1024*1024*20))
	if creds == nil {
		opts = append(opts, grpc.WithInsecure())
	} else {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	}
	opts = append(opts, dialOpts...)
	grpcConn, err := grpc.Dial(serverHost, opts...)
	return NewUserApiClient(grpcConn), err
}

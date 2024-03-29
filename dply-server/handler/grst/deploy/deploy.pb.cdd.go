// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: deploy.proto
package deploy

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
	DeployApi_DeployImage string
	DeployApi_Redeploy    string
}

var FullMethods = fullMethods{
	DeployApi_DeployImage: "/deploy.DeployApi/DeployImage",
	DeployApi_Redeploy:    "/deploy.DeployApi/Redeploy",
}

var NeedAuthFullMethods = []string{}

type AuthConfig struct {
	NeedAuth bool
	Roles    []string
}

var AuthConfigFullMethods = map[string]AuthConfig{
	"/deploy.DeployApi/DeployImage": AuthConfig{NeedAuth: false, Roles: []string{"*"}},
	"/deploy.DeployApi/Redeploy":    AuthConfig{NeedAuth: false, Roles: []string{"*"}},
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

/*==================== DeployApi Section ====================*/

func RegisterDeployApiGrstServer(grpcRestServer *grst.Server, hndl DeployApiServer) {

	forward_DeployApi_DeployImage_0 = grpcRestServer.GetForwardResponseMessage()

	forward_DeployApi_Redeploy_0 = grpcRestServer.GetForwardResponseMessage()

	RegisterDeployApiServer(grpcRestServer.GetGrpcServer(), hndl)
	grpcRestServer.RegisterRestHandler(RegisterDeployApiHandler)
}

func NewDeployApiGrstClient(serverHost string, creds *credentials.TransportCredentials, dialOpts ...grpc.DialOption) (DeployApiClient, error) {
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
	return NewDeployApiClient(grpcConn), err
}

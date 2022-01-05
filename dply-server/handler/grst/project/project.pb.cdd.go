// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: project.proto
package project

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
	ProjectApi_GetAll string
	ProjectApi_Create string
	ProjectApi_Delete string
}

var FullMethods = fullMethods{
	ProjectApi_GetAll: "/project.ProjectApi/GetAll",
	ProjectApi_Create: "/project.ProjectApi/Create",
	ProjectApi_Delete: "/project.ProjectApi/Delete",
}

var NeedAuthFullMethods = []string{}

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

/*==================== ProjectApi Section ====================*/

func RegisterProjectApiGrstServer(grpcRestServer *grst.Server, hndl ProjectApiServer) {

	forward_ProjectApi_GetAll_0 = grpcRestServer.GetForwardResponseMessage()

	forward_ProjectApi_Create_0 = grpcRestServer.GetForwardResponseMessage()

	forward_ProjectApi_Delete_0 = grpcRestServer.GetForwardResponseMessage()

	RegisterProjectApiServer(grpcRestServer.GetGrpcServer(), hndl)
	grpcRestServer.RegisterRestHandler(RegisterProjectApiHandler)
}

func NewProjectApiGrstClient(serverHost string, creds *credentials.TransportCredentials, dialOpts ...grpc.DialOption) (ProjectApiClient, error) {
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
	return NewProjectApiClient(grpcConn), err
}

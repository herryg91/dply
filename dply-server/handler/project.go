package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	project_usecase "github.com/herryg91/dply/dply-server/app/usecase/project"
	pbProject "github.com/herryg91/dply/dply-server/handler/grst/project"
	"google.golang.org/grpc/codes"
)

type handlerProject struct {
	pbProject.UnimplementedProjectApiServer
	project_uc project_usecase.UseCase
}

func NewProjectHandler(project_uc project_usecase.UseCase) pbProject.ProjectApiServer {
	return &handlerProject{project_uc: project_uc}
}

func (h *handlerProject) GetAll(ctx context.Context, req *empty.Empty) (*pbProject.Projects, error) {
	projects, err := h.project_uc.Get()
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10000, err.Error())
	}

	resp := &pbProject.Projects{
		Projects: []*pbProject.Project{},
	}
	for _, p := range projects {
		resp.Projects = append(resp.Projects, &pbProject.Project{
			Id:          int32(p.Id),
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return resp, nil
}

func (h *handlerProject) Create(ctx context.Context, req *pbProject.CreateReq) (*empty.Empty, error) {
	if err := pbProject.ValidateRequest(req); err != nil {
		return nil, err
	}
	err := h.project_uc.Create(req.Name, req.Description)
	if err != nil {
		if errors.Is(err, project_usecase.ErrProjectNameInvalidFormat) {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 10001, err.Error(), &grst_errors.ErrorDetail{
				Code:    1,
				Field:   "name",
				Message: err.Error(),
			})
		} else if errors.Is(err, project_usecase.ErrProjectAlreadyExist) {
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10002, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10000, err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *handlerProject) Delete(ctx context.Context, req *pbProject.DeleteReq) (*empty.Empty, error) {
	if err := pbProject.ValidateRequest(req); err != nil {
		return nil, err
	}
	err := h.project_uc.Delete(req.Name)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10000, err.Error())
	}

	return &empty.Empty{}, nil
}

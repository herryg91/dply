package project_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbProject "github.com/herryg91/dply/dply/clients/grst/project"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
)

type repository struct {
	cli pbProject.ProjectApiClient
}

func New(cli pbProject.ProjectApiClient) repository_intf.ProjectRepository {
	return &repository{cli: cli}
}

func (r *repository) GetAll() ([]entity.Project, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return []entity.Project{}, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	data, err := r.cli.GetAll(ctx, &empty.Empty{})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return []entity.Project{}, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return []entity.Project{}, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return []entity.Project{}, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	resp := []entity.Project{}
	for _, p := range data.Projects {
		resp = append(resp, entity.Project{
			Id:          int(p.Id),
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return resp, nil
}
func (r *repository) Create(p entity.Project) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.Create(ctx, &pbProject.CreateReq{Name: p.Name, Description: p.Description})
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

func (r *repository) Delete(name string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.Delete(ctx, &pbProject.DeleteReq{Name: name})
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

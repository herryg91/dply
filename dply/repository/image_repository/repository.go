package image_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbImage "github.com/herryg91/dply/dply/clients/grst/image"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
)

type repository struct {
	cli pbImage.ImageApiClient
}

func New(cli pbImage.ImageApiClient) repository_intf.ImageRepository {
	return &repository{cli}
}

func (r *repository) Add(repoName, image, description string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.Add(ctx, &pbImage.AddReq{
		Image:       image,
		Repository:  repoName,
		Description: description,
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) Remove(repoName, digest string) error {
	// Notes: unimplementated for now
	return nil
}
func (r *repository) Get(repoName string, page, size int) ([]entity.ContainerImage, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return []entity.ContainerImage{}, fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	datas, err := r.cli.Get(ctx, &pbImage.GetReq{
		Repository: repoName,
		Page:       int32(page),
		Size:       int32(size),
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return []entity.ContainerImage{}, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return []entity.ContainerImage{}, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return []entity.ContainerImage{}, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	resp := []entity.ContainerImage{}

	for _, data := range datas.Images {
		createdAt := data.CreatedAt.AsTime()
		resp = append(resp, entity.ContainerImage{
			Id:             int(data.Id),
			Digest:         data.Digest,
			Image:          data.Image,
			RepositoryName: data.Repository,
			Description:    data.Description,
			CreatedBy:      int(data.CreatedBy),
			CreatedAt:      &createdAt,
		})
	}
	return resp, nil
}

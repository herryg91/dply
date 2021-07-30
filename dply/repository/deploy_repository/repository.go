package deploy_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbDeploy "github.com/herryg91/dply/dply/clients/grst/deploy"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
)

type repository struct {
	cli pbDeploy.DeployApiClient
}

func New(cli pbDeploy.DeployApiClient) repository_intf.DeployRepository {
	return &repository{cli}
}

func (r *repository) Deploy(env, name, digest string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.DeployImage(ctx, &pbDeploy.DeployImageReq{Env: env, Name: name, Digest: digest})
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

func (r *repository) Redeploy(env, name string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, "You are not login")
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.Redeploy(ctx, &pbDeploy.RedeployReq{Env: env, Name: name})
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

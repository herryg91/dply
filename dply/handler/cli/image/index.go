package cli_image

import (
	"log"

	"github.com/herryg91/dply/dply/app/repository"
	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	pbImage "github.com/herryg91/dply/dply/clients/grst/image"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/image_repository"
	"github.com/spf13/cobra"
)

type CmdImage struct {
	*cobra.Command
}

func New() *CmdImage {
	c := &CmdImage{}
	c.Command = &cobra.Command{
		Use:   "image",
		Short: "Container image management",
		Long:  "Container image management",
	}

	cfg := entity.Config{}.FromFile()
	var image_repo repository.ImageRepository = nil
	var image_uc image_usecase.UseCase = nil
	var imageCli pbImage.ImageApiClient = nil
	if cfg != nil {
		var err error
		imageCli, err = pbImage.NewImageApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		image_repo, err = image_repository.New(imageCli, cfg)
		if err != nil {
			log.Panicln("Failed to initialized image repository", err)
		}
		image_uc = image_usecase.New(image_repo)
	}
	c.AddCommand(newCmdImageAdd(cfg, image_uc).Command)
	c.AddCommand(newCmdImageList(cfg, image_uc).Command)
	c.AddCommand(newCmdImageCreate(cfg, image_uc).Command)
	return c
}

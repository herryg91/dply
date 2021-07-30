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

	setting := entity.Setting{}.FromFile()
	var image_repo repository.ImageRepository = nil
	var image_uc image_usecase.UseCase = nil
	var imageCli pbImage.ImageApiClient = nil
	if setting != nil {
		var err error
		imageCli, err = pbImage.NewImageApiGrstClient(setting.ServerHostGrpc, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		image_repo = image_repository.New(imageCli)
		image_uc = image_usecase.New(image_repo)
	}
	c.AddCommand(newCmdImageAdd(image_uc).Command)
	c.AddCommand(newCmdImageList(image_uc).Command)
	// c.AddCommand(newCmdImageRemove(image_uc).Command)
	return c
}

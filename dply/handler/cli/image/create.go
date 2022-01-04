package cli_image

import (
	"errors"
	"log"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"

	"github.com/spf13/cobra"
)

type CmdImageCreate struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	name        string
	description string
}

func newCmdImageCreate(image_uc image_usecase.UseCase) *CmdImageCreate {
	c := &CmdImageCreate{image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "create",
		Short: "Create image",
		Long:  "Create image",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "service (repository) name of image")
	c.Command.Flags().StringVarP(&c.description, "desc", "d", "", "image description")
	return c
}

func (c *CmdImageCreate) runCommand(cmd *cobra.Command, args []string) error {
	if c.image_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.description == "" {
		log.Println("dasndjsanjdkasn")
		return errors.New("`--desc / -d` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	cfg := entity.Config{}.FromFile()
	err := c.image_uc.Create(c.name, cfg.RegistryTagPrefix, c.description)
	if err != nil {
		return err
	}

	return nil
}

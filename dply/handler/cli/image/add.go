package cli_image

import (
	"errors"
	"fmt"
	"strings"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"

	"github.com/spf13/cobra"
)

type CmdImageAdd struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	name        string
	image       string
	description string
}

func newCmdImageAdd(image_uc image_usecase.UseCase) *CmdImageAdd {
	c := &CmdImageAdd{image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "add",
		Short: "Add container image",
		Long:  "Add container image",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "service (repository) name of image")
	c.Command.Flags().StringVarP(&c.image, "image", "i", "", "full image path with digest format (<repo_name>@<digest>). Example: gcr.io/repo@sha256:xxx")
	c.Command.Flags().StringVarP(&c.description, "description", "d", "", "image description")
	return c
}

func (c *CmdImageAdd) runCommand(cmd *cobra.Command, args []string) error {
	if c.image_uc == nil {
		return errors.New("You haven't configure setting. command: `dply-cli setting --server=<dply_server_host>`")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("parameter --name or -n is required")
		}
		c.name = data.Name
	} else if c.image == "" {
		return errors.New("parameter --image or -i is required")
	} else if c.image != "" {
		imageSplit := strings.Split(c.image, "@")
		if len(imageSplit) != 2 {
			return errors.New("parameter --image or -i has invalid format. Expected format: <repo_name>@<digest>. Example: gcr.io/repo@sha256:xxx ")
		}
	}

	err := c.image_uc.Add(c.name, c.image, c.description)
	if err != nil {
		return err
	}
	fmt.Println("image `" + c.image + " succesfully added")

	return nil
}

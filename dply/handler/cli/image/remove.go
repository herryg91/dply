package cli_image

import (
	"errors"
	"fmt"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"

	"github.com/spf13/cobra"
)

type CmdImageRemove struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	serviceName string
	digest      string
}

func newCmdImageRemove(image_uc image_usecase.UseCase) *CmdImageRemove {
	c := &CmdImageRemove{image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "remove",
		Short: "Remove container image (require admin access)",
		Long:  "Remove container image (require admin access)",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.serviceName, "service", "s", "", "service (repository) name of image")
	c.Command.Flags().StringVarP(&c.digest, "digest", "d", "", "image digest")
	return c
}

func (c *CmdImageRemove) runCommand(cmd *cobra.Command, args []string) error {
	if c.image_uc == nil {
		return errors.New("You haven't init config. command: `dply init config --server=<dply_server_host> --name=<name> --email=<email>`")
	} else if c.serviceName == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`-s|--service` is required or you are in directory with service.yaml (has name variable)")
		}
		c.serviceName = data.Name
	} else if c.digest == "" {
		return errors.New("digest (-d|--digest) is required")
	}

	err := c.image_uc.Remove(c.serviceName, c.digest)
	if err != nil {
		return err
	}
	fmt.Println("" + c.digest + " succesfully removed")

	return nil
}

package cli_envar

import (
	"errors"
	"fmt"

	deploy_usecase "github.com/herryg91/dply/dply/app/usecase/deploy"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/spf13/cobra"
)

type CmdDeployImage struct {
	*cobra.Command

	deploy_uc deploy_usecase.UseCase
	env       string
	name      string
}

func newDeployImage(deploy_uc deploy_usecase.UseCase) *CmdDeployImage {
	c := &CmdDeployImage{
		deploy_uc: deploy_uc,
	}
	c.Command = &cobra.Command{
		Use:   "image",
		Short: "deploy image <digest> -n <name> -e <environment/namespace>",
		Long:  "deploy image <digest> -n <name> -e <environment/namespace>",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment / namespace of service")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdDeployImage) runCommand(cmd *cobra.Command, args []string) error {
	if c.deploy_uc == nil {
		return errors.New("You haven't configure setting. command: `dply-cli setting --server=<dply_server_host>`")
	} else if len(args) <= 0 {
		return errors.New("deploy image <digest>. 'digest' parameter is required")
	} else if c.env == "" {
		return errors.New("`-e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("param --name or -n is required")
		}
		c.name = data.Name
	}
	digest := args[0]

	err := c.deploy_uc.Deploy(c.env, c.name, digest)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Deploying %s (%s) with digest %s. To monitor the process: kubectl get pod -n %s -lapp=%s", c.env, c.name, digest, c.env, c.name))

	return nil
}

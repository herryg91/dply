package cli_envar

import (
	"errors"
	"fmt"

	deploy_usecase "github.com/herryg91/dply/dply/app/usecase/deploy"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/spf13/cobra"
)

type CmdDeployReload struct {
	*cobra.Command
	deploy_uc deploy_usecase.UseCase

	project string
	env     string
	name    string
}

func newDeployReload(cfg *entity.Config, deploy_uc deploy_usecase.UseCase) *CmdDeployReload {
	c := &CmdDeployReload{project: cfg.Project, deploy_uc: deploy_uc}
	c.Command = &cobra.Command{
		Use:   "reload",
		Short: "reload deployment",
		Long:  "reload deployment",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment / namespace of service")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdDeployReload) runCommand(cmd *cobra.Command, args []string) error {
	service_yaml_data, err := serviceYaml.GetServiceYAML("service.yaml")
	if err == nil {
		if service_yaml_data.Project != "" {
			c.project = service_yaml_data.Project
		}
	}

	if c.deploy_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.env == "" {
		return errors.New("`-e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("param --name or -n is required")
		}
		c.name = data.Name
	}

	err = c.deploy_uc.Redeploy(c.project, c.env, c.name)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Redeploying %s (%s). To monitor the process: kubectl get pod -n %s -lapp=%s", c.env, c.name, c.env, c.name))

	return nil
}

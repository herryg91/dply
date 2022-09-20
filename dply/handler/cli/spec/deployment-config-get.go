package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	deployment_config_usecase "github.com/herryg91/dply/dply/app/usecase/deployment-config"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecDeploymentConfigGet struct {
	*cobra.Command
	deployment_config_uc deployment_config_usecase.UseCase

	project string
	env     string
	name    string
}

func newSpecDeploymentConfigGet(cfg *entity.Config, deployment_config_uc deployment_config_usecase.UseCase) *CmdSpecDeploymentConfigGet {
	c := &CmdSpecDeploymentConfigGet{project: cfg.Project, deployment_config_uc: deployment_config_uc}
	c.Command = &cobra.Command{
		Use:     "deployment-config-get",
		Aliases: []string{"dcg"},
		Short:   "View deployment configuration specification",
		Long:    "View deployment configuration specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecDeploymentConfigGet) runCommand(cmd *cobra.Command, args []string) error {
	if c.deployment_config_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.env == "" {
		return errors.New("`--env / -e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	resp, err := c.deployment_config_uc.Get(c.project, c.env, c.name)
	if err != nil {
		if errors.Is(err, deployment_config_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	jsonData, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

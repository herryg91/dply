package cli_spec

import (
	"errors"
	"fmt"

	deployment_config_usecase "github.com/herryg91/dply/dply/app/usecase/deployment-config"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecDeploymentConfigEdit struct {
	*cobra.Command
	deployment_config_uc deployment_config_usecase.UseCase

	project string
	env     string
	name    string
	editor  string
}

func newSpecDeploymentConfigEdit(cfg *entity.Config, deployment_config_uc deployment_config_usecase.UseCase) *CmdSpecDeploymentConfigEdit {
	c := &CmdSpecDeploymentConfigEdit{
		project:              cfg.Project,
		deployment_config_uc: deployment_config_uc,
		editor:               cfg.Editor,
	}
	c.Command = &cobra.Command{
		Use:     "deployment-config-edit",
		Aliases: []string{"dce"},
		Short:   "Edit deployment configuration specification",
		Long:    "Edit deployment configuration specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecDeploymentConfigEdit) runCommand(cmd *cobra.Command, args []string) error {
	if c.deployment_config_uc == nil {
		return errors.New("You haven't configure config. command: `dply-cli config --server=<dply_server_host>`")
	} else if c.env == "" {
		return errors.New("`-e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	ok, err := c.deployment_config_uc.UpsertViaEditor(c.project, c.env, c.name, editor.EditorApp(c.editor))
	if err != nil {
		if errors.Is(err, deployment_config_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}
	if ok {
		fmt.Println("port specification " + c.name + " (" + c.env + ") succesfully updated")
	} else {
		fmt.Println("Nothing to change")
	}

	return nil
}

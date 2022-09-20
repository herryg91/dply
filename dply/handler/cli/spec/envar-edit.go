package cli_spec

import (
	"errors"
	"fmt"

	envar_usecase "github.com/herryg91/dply/dply/app/usecase/envar"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecEnvarEdit struct {
	*cobra.Command
	envar_uc envar_usecase.UseCase

	project string
	env     string
	name    string
	editor  string
}

func newSpecEnvarEdit(cfg *entity.Config, envar_uc envar_usecase.UseCase) *CmdSpecEnvarEdit {
	c := &CmdSpecEnvarEdit{
		envar_uc: envar_uc,
		project:  cfg.Project,
		editor:   cfg.Editor,
	}
	c.Command = &cobra.Command{
		Use:     "envar-edit",
		Aliases: []string{"ee"},
		Short:   "Edit environment variables",
		Long:    "Edit environment variables",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecEnvarEdit) runCommand(cmd *cobra.Command, args []string) error {
	service_yaml_data, err := serviceYaml.GetServiceYAML("service.yaml")
	if err != nil {
		return err
	}
	if service_yaml_data.Project != "" {
		c.project = service_yaml_data.Project
	}

	if c.envar_uc == nil {
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

	ok, err := c.envar_uc.UpsertViaEditor(c.project, c.env, c.name, editor.EditorApp(c.editor))
	if err != nil {
		if errors.Is(err, envar_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}
	if ok {
		fmt.Println("environment variable " + c.name + " (" + c.env + ") succesfully updated")
	} else {
		fmt.Println("Nothing to change")
	}

	return nil
}

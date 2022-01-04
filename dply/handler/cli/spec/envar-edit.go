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
	cfg      *entity.Config

	env  string
	name string
}

func newSpecEnvarEdit(envar_uc envar_usecase.UseCase, cfg *entity.Config) *CmdSpecEnvarEdit {
	c := &CmdSpecEnvarEdit{envar_uc: envar_uc, cfg: cfg}
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
	if c.envar_uc == nil {
		return errors.New("You haven't configure config. command: `dply-cli config --server=<dply_server_host>`")
	} else if c.cfg == nil {
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

	ok, err := c.envar_uc.UpsertViaEditor(c.env, c.name, editor.EditorApp(c.cfg.Editor))
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

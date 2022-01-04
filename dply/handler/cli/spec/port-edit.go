package cli_spec

import (
	"errors"
	"fmt"

	port_usecase "github.com/herryg91/dply/dply/app/usecase/port"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecPortEdit struct {
	*cobra.Command
	port_uc port_usecase.UseCase
	cfg     *entity.Config

	env  string
	name string
}

func newSpecPortEdit(port_uc port_usecase.UseCase, cfg *entity.Config) *CmdSpecPortEdit {
	c := &CmdSpecPortEdit{port_uc: port_uc, cfg: cfg}
	c.Command = &cobra.Command{
		Use:     "port-edit",
		Aliases: []string{"pe"},
		Short:   "Edit port specification",
		Long:    "Edit port specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecPortEdit) runCommand(cmd *cobra.Command, args []string) error {
	if c.port_uc == nil {
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

	ok, err := c.port_uc.UpsertViaEditor(c.env, c.name, editor.EditorApp(c.cfg.Editor))
	if err != nil {
		if errors.Is(err, port_usecase.ErrUnauthorized) {
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

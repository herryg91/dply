package cli_spec

import (
	"errors"
	"fmt"

	scale_usecase "github.com/herryg91/dply/dply/app/usecase/scale"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecScalingEdit struct {
	*cobra.Command
	scale_uc scale_usecase.UseCase
	cfg      *entity.Config

	project string
	env     string
	name    string
	editor  string
}

func newSpecScalingEdit(cfg *entity.Config, scale_uc scale_usecase.UseCase) *CmdSpecScalingEdit {
	c := &CmdSpecScalingEdit{
		project:  cfg.Project,
		scale_uc: scale_uc, cfg: cfg,
		editor: cfg.Editor,
	}
	c.Command = &cobra.Command{
		Use:     "scaling-edit",
		Aliases: []string{"se"},
		Short:   "Edit scaling specification",
		Long:    "Edit scaling specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecScalingEdit) runCommand(cmd *cobra.Command, args []string) error {
	if c.scale_uc == nil {
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

	ok, err := c.scale_uc.UpsertViaEditor(c.project, c.env, c.name, editor.EditorApp(c.editor))
	if err != nil {
		if errors.Is(err, scale_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}
	if ok {
		fmt.Println("scaling config " + c.name + " (" + c.env + ") succesfully updated")
	} else {
		fmt.Println("Nothing to change")
	}

	return nil
}

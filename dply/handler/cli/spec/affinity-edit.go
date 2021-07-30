package cli_spec

import (
	"errors"
	"fmt"

	affinity_usecase "github.com/herryg91/dply/dply/app/usecase/affinity"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/editor"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecAffinityEdit struct {
	*cobra.Command
	affinity_uc affinity_usecase.UseCase
	setting     *entity.Setting

	env  string
	name string
}

func newSpecAffinityEdit(affinity_uc affinity_usecase.UseCase, setting *entity.Setting) *CmdSpecAffinityEdit {
	c := &CmdSpecAffinityEdit{affinity_uc: affinity_uc, setting: setting}
	c.Command = &cobra.Command{
		Use:     "affinity-edit",
		Aliases: []string{"ae"},
		Short:   "Edit affinity specification",
		Long:    "Edit affinity specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecAffinityEdit) runCommand(cmd *cobra.Command, args []string) error {
	if c.affinity_uc == nil {
		return errors.New("You haven't configure setting. command: `dply-cli setting --server=<dply_server_host>`")
	} else if c.setting == nil {
		return errors.New("You haven't configure setting. command: `dply-cli setting --server=<dply_server_host>`")
	} else if c.env == "" {
		return errors.New("`-e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	ok, err := c.affinity_uc.UpsertViaEditor(c.env, c.name, editor.EditorApp(c.setting.Editor))
	if err != nil {
		if errors.Is(err, affinity_usecase.ErrUnauthorized) {
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

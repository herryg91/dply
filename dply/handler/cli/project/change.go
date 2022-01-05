package cli_project

import (
	"errors"

	"github.com/herryg91/dply/dply/entity"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type CmdProjectChange struct {
	*cobra.Command
	project string
}

func newCmdProjectChange() *CmdProjectChange {
	c := &CmdProjectChange{}
	c.Command = &cobra.Command{
		Use:   "change",
		Short: "Change project",
		Long:  "Change project",
	}
	c.RunE = c.runCommand
	// c.Command.Flags().StringVarP(&c.project, "project", "p", "", "Project Name")
	return c
}

func (c *CmdProjectChange) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("dply projec change <name>. Name is required")
	}
	c.project = args[0]

	cfg := entity.Config{}.FromFile()
	err := cfg.UpdateProject(c.project)
	if err != nil {
		return err
	}

	logrus.Infoln("Project changed into: " + c.project)
	return nil
}

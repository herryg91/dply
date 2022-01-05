package cli_project

import (
	"errors"
	"fmt"

	project_usecase "github.com/herryg91/dply/dply/app/usecase/project"
	"github.com/spf13/cobra"
)

type CmdProjectDelete struct {
	*cobra.Command
	project_uc project_usecase.UseCase
	name       string
}

func newCmdProjectDelete(project_uc project_usecase.UseCase) *CmdProjectDelete {
	c := &CmdProjectDelete{project_uc: project_uc}
	c.Command = &cobra.Command{
		Use:   "delete",
		Short: "Delete project",
		Long:  "Delete project",
	}
	c.RunE = c.runCommand

	return c
}

func (c *CmdProjectDelete) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("dply projec change <name>. Name is required")
	}
	c.name = args[0]

	err := c.project_uc.Delete(c.name)
	if err != nil {
		return err
	}
	fmt.Println("Project is deleted: " + c.name)
	return nil
}

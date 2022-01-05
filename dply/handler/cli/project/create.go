package cli_project

import (
	"fmt"

	project_usecase "github.com/herryg91/dply/dply/app/usecase/project"
	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdProjectCreate struct {
	*cobra.Command
	project_uc  project_usecase.UseCase
	name        string
	description string
}

func newCmdProjectCreate(project_uc project_usecase.UseCase) *CmdProjectCreate {
	c := &CmdProjectCreate{project_uc: project_uc}
	c.Command = &cobra.Command{
		Use:   "create",
		Short: "Create new project",
		Long:  "Create new project",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "project name (required)")
	c.Command.Flags().StringVarP(&c.description, "desc", "d", "", "project description")

	return c
}

func (c *CmdProjectCreate) runCommand(cmd *cobra.Command, args []string) error {
	err := c.project_uc.Create(entity.Project{Name: c.name, Description: c.description})
	if err != nil {
		return err
	}
	fmt.Println("Project is created: " + c.name)
	return nil
}

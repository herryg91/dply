package cli_project

import (
	"fmt"

	project_usecase "github.com/herryg91/dply/dply/app/usecase/project"
	"github.com/spf13/cobra"
)

type CmdProjectList struct {
	*cobra.Command
	project_uc project_usecase.UseCase
}

func newCmdProjectList(project_uc project_usecase.UseCase) *CmdProjectList {
	c := &CmdProjectList{project_uc: project_uc}
	c.Command = &cobra.Command{
		Use:   "list",
		Short: "Print all available project",
		Long:  "Print all available project",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdProjectList) runCommand(cmd *cobra.Command, args []string) error {
	fmt.Println("Project List")
	projects, err := c.project_uc.Get()
	if err != nil {
		return err
	}
	for idx, p := range projects {
		fmt.Println(fmt.Sprintf("%d. %s", idx+1, p.Name))
	}
	return nil
}

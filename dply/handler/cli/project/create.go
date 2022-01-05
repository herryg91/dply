package cli_project

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CmdProjectCreate struct {
	*cobra.Command
}

func newCmdProjectCreate() *CmdProjectCreate {
	c := &CmdProjectCreate{}
	c.Command = &cobra.Command{
		Use:   "create",
		Short: "Create new project",
		Long:  "Create new project",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdProjectCreate) runCommand(cmd *cobra.Command, args []string) error {
	fmt.Println("Under development")
	return nil
}

package cli_project

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CmdProjectList struct {
	*cobra.Command
}

func newCmdProjectList() *CmdProjectList {
	c := &CmdProjectList{}
	c.Command = &cobra.Command{
		Use:   "list",
		Short: "Print all available project",
		Long:  "Print all available project",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdProjectList) runCommand(cmd *cobra.Command, args []string) error {
	fmt.Println("Under development")
	return nil
}

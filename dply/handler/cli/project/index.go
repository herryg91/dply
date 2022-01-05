package cli_project

import (
	"github.com/spf13/cobra"
)

type CmdConfig struct {
	*cobra.Command
}

func New() *CmdConfig {
	c := &CmdConfig{}
	c.Command = &cobra.Command{
		Use:   "project",
		Short: "Configure project",
		Long:  "Configure project",
	}

	c.AddCommand(newCmdProjectChange().Command)
	c.AddCommand(newCmdProjectList().Command)
	c.AddCommand(newCmdProjectCreate().Command)

	return c
}

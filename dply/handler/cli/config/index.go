package cli_config

import (
	"github.com/spf13/cobra"
)

type CmdConfig struct {
	*cobra.Command
}

func New() *CmdConfig {
	c := &CmdConfig{}
	c.Command = &cobra.Command{
		Use:   "config",
		Short: "Configure server & preferences",
		Long:  "Configure server & preferences",
	}

	c.AddCommand(newConfigEdit().Command)
	c.AddCommand(newConfigGet().Command)
	c.AddCommand(newConfigSetDplyServer().Command)
	c.AddCommand(newConfigSetDocker().Command)
	c.AddCommand(newConfigSetDockerCert().Command)
	c.AddCommand(newConfigSetRegistryHost().Command)
	c.AddCommand(newConfigSetRegistryUsername().Command)
	c.AddCommand(newConfigSetRegistryPassword().Command)
	c.AddCommand(newConfigSetEditor().Command)
	return c
}

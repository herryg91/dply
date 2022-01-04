package cli_config

import (
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetRegistryHost struct {
	*cobra.Command
}

func newConfigSetRegistryHost() *CmdConfigSetRegistryHost {
	c := &CmdConfigSetRegistryHost{}
	c.Command = &cobra.Command{
		Use:   "set-registry-host",
		Short: "Set registry server (ex: https://registry.hub.docker.com ,https://gcr.io, etc)",
		Long:  "Set registry server (ex: https://registry.hub.docker.com ,https://gcr.io, etc)",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetRegistryHost) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require value parameter")
	}

	new_registry_host := args[0]
	cfg := entity.Config{}.FromFile()
	old_registry_host := cfg.RegistryHost
	if cfg.RegistryHost != new_registry_host {
		cfg.RegistryHost = new_registry_host
		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Registry server host was changed, %s -> %s", old_registry_host, new_registry_host))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

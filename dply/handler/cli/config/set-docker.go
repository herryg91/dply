package cli_config

import (
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetDocker struct {
	*cobra.Command
}

func newConfigSetDocker() *CmdConfigSetDocker {
	c := &CmdConfigSetDocker{}
	c.Command = &cobra.Command{
		Use:   "set-docker",
		Short: "Set docker engine server (tcp://...)",
		Long:  "Set docker engine server (tcp://...)",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetDocker) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require value parameter")
	}

	new_docker_host := args[0]
	cfg := entity.Config{}.FromFile()
	old_docker_host := cfg.DockerHost
	if cfg.DockerHost != new_docker_host {
		cfg.DockerHost = new_docker_host
		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Docker engine host was changed, %s -> %s", old_docker_host, new_docker_host))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

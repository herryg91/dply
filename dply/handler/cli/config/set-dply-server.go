package cli_config

import (
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetDplyServer struct {
	*cobra.Command
}

func newConfigSetDplyServer() *CmdConfigSetDplyServer {
	c := &CmdConfigSetDplyServer{}
	c.Command = &cobra.Command{
		Use:   "set-dply-server",
		Short: "Set dply server host",
		Long:  "Set dply server host",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetDplyServer) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require value parameter")
	}

	new_server_host := args[0]
	cfg := entity.Config{}.FromFile()
	old_server_host := cfg.DplyServerHost
	if cfg.DplyServerHost != new_server_host {
		cfg.DplyServerHost = new_server_host
		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Dply Server Host was changed, %s -> %s", old_server_host, new_server_host))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

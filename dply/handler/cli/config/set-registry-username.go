package cli_config

import (
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetRegistryUsername struct {
	*cobra.Command
}

func newConfigSetRegistryUsername() *CmdConfigSetRegistryUsername {
	c := &CmdConfigSetRegistryUsername{}
	c.Command = &cobra.Command{
		Use:   "set-registry-username",
		Short: "Set registry username",
		Long:  "Set registry username",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetRegistryUsername) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require value parameter")
	}

	new_registry_username := args[0]
	cfg := entity.Config{}.FromFile()
	old_registry_username := cfg.RegistryUsername
	if cfg.RegistryUsername != new_registry_username {
		cfg.RegistryUsername = new_registry_username
		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Registry server username was changed, %s -> %s", old_registry_username, new_registry_username))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

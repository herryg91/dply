package cli_config

import (
	"fmt"
	"io/ioutil"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetRegistryPassword struct {
	*cobra.Command
	from_file string
}

func newConfigSetRegistryPassword() *CmdConfigSetRegistryPassword {
	c := &CmdConfigSetRegistryPassword{}
	c.Command = &cobra.Command{
		Use:   "set-registry-password",
		Short: "Set registry password",
		Long:  "Set registry password",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.from_file, "from-file", "f", "", "password file location")

	return c
}

func (c *CmdConfigSetRegistryPassword) runCommand(cmd *cobra.Command, args []string) error {
	password := ""
	if c.from_file != "" {
		from_file_content, err := ioutil.ReadFile(c.from_file)
		if err != nil {
			return err
		}
		password = string(from_file_content)
	} else {
		if len(args) == 0 {
			return fmt.Errorf("Parameter password is required. `dply config set-registry-password <password>`")
		}
		password = args[0]
	}

	cfg := entity.Config{}.FromFile()
	if cfg.RegistryPassword == password {
		fmt.Println("Nothing to update")
		return nil
	}
	cfg.RegistryPassword = password
	err := cfg.SaveConfig()
	if err != nil {
		return err
	}
	fmt.Println("Registry password successfully updated")

	return nil
}

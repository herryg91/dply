package cli_config

import (
	"encoding/json"
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigGet struct {
	*cobra.Command
}

func newConfigGet() *CmdConfigGet {
	c := &CmdConfigGet{}
	c.Command = &cobra.Command{
		Use:   "get",
		Short: "Print current configuration & preferences",
		Long:  "Print current configuration & preferences",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigGet) runCommand(cmd *cobra.Command, args []string) error {
	cfg := entity.Config{}.FromFile()
	current_data, _ := json.MarshalIndent(cfg, "", "    ")

	fmt.Println(string(current_data))
	return nil
}

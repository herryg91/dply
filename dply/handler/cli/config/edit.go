package cli_config

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigEdit struct {
	*cobra.Command
}

func newConfigEdit() *CmdConfigEdit {
	c := &CmdConfigEdit{}
	c.Command = &cobra.Command{
		Use:   "edit",
		Short: "Edit dply configuration & preferences",
		Long:  "Edit dply configuration & preferences",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigEdit) runCommand(cmd *cobra.Command, args []string) error {
	cfg := entity.Config{}.FromFile()

	err := cfg.UpdateByEditor()
	if err != nil {
		if errors.Is(err, entity.ErrConfigNothingChange) {
			fmt.Println("Nothing to change")
			return nil
		}
		return err
	}

	fmt.Println("Config file successfully saved")
	return nil
}

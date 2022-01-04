package cli_config

import (
	"fmt"
	"strings"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetEditor struct {
	*cobra.Command
}

func newConfigSetEditor() *CmdConfigSetEditor {
	c := &CmdConfigSetEditor{}
	c.Command = &cobra.Command{
		Use:   "set-editor",
		Short: "Set terminal editor. Choice: vi|nano",
		Long:  "Set terminal editor. Choice: vi|nano",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetEditor) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require parameter, choice: vi|nano")
	}
	new_editor := args[0]
	if strings.ToLower(new_editor) != "vi" && strings.ToLower(new_editor) != "nano" {
		return fmt.Errorf("invalid parameter, choice: vi|nano")
	}

	cfg := entity.Config{}.FromFile()

	if new_editor != cfg.Editor {
		old_editor := cfg.Editor
		cfg.Editor = new_editor

		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Editor was changed, %s -> %s", old_editor, new_editor))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

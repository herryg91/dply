package cli

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdSetting struct {
	*cobra.Command
	serverHostGrpc string
	editor         string
}

func NewCmdSetting() *CmdSetting {
	c := &CmdSetting{}
	c.Command = &cobra.Command{
		Use:   "setting",
		Short: "Configure dply-server grpc host & editor",
		Long:  "Configure dply-server grpc host & editor",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.serverHostGrpc, "server", "s", "", "dply-server grpc host. example: dply-server.dply.svc.cluster.local:9090")
	c.Command.Flags().StringVarP(&c.editor, "editor", "e", "", "Text editor app in your computer, example: vi|nano")

	return c
}

func (c *CmdSetting) runCommand(cmd *cobra.Command, args []string) error {
	if c.serverHostGrpc == "" && c.editor == "" {
		return errors.New("require flags / parameter")
	}

	isUpdate := false
	setting := entity.Setting{}.FromFile()
	if c.serverHostGrpc != "" && c.serverHostGrpc != setting.ServerHostGrpc {
		isUpdate = true
		setting.ServerHostGrpc = c.serverHostGrpc
	}
	if c.editor != "" && c.editor != setting.Editor {
		isUpdate = true
		setting.Editor = c.editor
	}
	if isUpdate {
		setting.SaveSetting()
		fmt.Println("setting.json is updated.")
		fmt.Println("dply-server Host (GRPC) : " + setting.ServerHostGrpc)
		fmt.Println("Editor                  : " + setting.Editor)
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

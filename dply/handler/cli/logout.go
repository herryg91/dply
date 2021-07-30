package cli

import (
	auth_usecase "github.com/herryg91/dply/dply/app/usecase/auth"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type CmdLogout struct {
	*cobra.Command
	auth_uc auth_usecase.UseCase
}

func NewCmdLogout() *CmdLogout {
	c := &CmdLogout{
		auth_uc: auth_usecase.New(nil),
	}
	c.Command = &cobra.Command{
		Use:   "logout",
		Short: "logout",
		Long:  "logout",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdLogout) runCommand(cmd *cobra.Command, args []string) error {
	c.auth_uc.Logout()
	logrus.Infoln("You are logout")
	return nil
}

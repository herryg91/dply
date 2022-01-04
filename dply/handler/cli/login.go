package cli

import (
	"errors"
	"log"

	"github.com/badoux/checkmail"
	"github.com/herryg91/dply/dply/app/repository"
	auth_usecase "github.com/herryg91/dply/dply/app/usecase/auth"
	pbUser "github.com/herryg91/dply/dply/clients/grst/user"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/user_repository"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type CmdLogin struct {
	*cobra.Command
	auth_uc auth_usecase.UseCase

	email    string
	password string
}

func NewCmdLogin() *CmdLogin {
	cfg := entity.Config{}.FromFile()
	var user_repo repository.UserRepository = nil
	var auth_uc auth_usecase.UseCase = nil
	var userCli pbUser.UserApiClient = nil
	if cfg != nil {
		var err error
		userCli, err = pbUser.NewUserApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server:", err)
		}

		user_repo = user_repository.New(userCli)
		auth_uc = auth_usecase.New(user_repo)
	}
	c := &CmdLogin{
		auth_uc: auth_uc,
	}
	c.Command = &cobra.Command{
		Use:   "login",
		Short: "authentication to dply-server",
		Long:  "authentication to dply-server",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.email, "email", "e", "", "Enter your email")
	c.Command.Flags().StringVarP(&c.password, "password", "p", "", "Enter your password")
	return c
}

func (c *CmdLogin) runCommand(cmd *cobra.Command, args []string) error {
	if c.password == "" {
		return errors.New("`--password / -p` is required")
	} else if c.email == "" {
		return errors.New("`--email / -e` is required")
	} else if err := checkmail.ValidateFormat(c.email); err != nil {
		return errors.New("`--email / -e` is not email format, got: " + c.email)
	} else if c.auth_uc == nil {
		return errors.New("You haven't configure dply-server host. Run `dply config edit`")
	}

	err := c.auth_uc.Login(c.email, c.password)
	if err != nil {
		if errors.Is(err, auth_usecase.ErrLoginFailed) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}
	logrus.Infoln("Login success")
	return nil
}

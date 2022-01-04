package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/herryg91/dply/dply/app/repository"
	auth_usecase "github.com/herryg91/dply/dply/app/usecase/auth"
	server_usecase "github.com/herryg91/dply/dply/app/usecase/server"
	pbServer "github.com/herryg91/dply/dply/clients/grst/server"
	pbUser "github.com/herryg91/dply/dply/clients/grst/user"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/server_repository"
	"github.com/herryg91/dply/dply/repository/user_repository"
	"github.com/spf13/cobra"
)

type CmdStatus struct {
	*cobra.Command
	auth_uc   auth_usecase.UseCase
	server_uc server_usecase.UseCase
}

func NewCmdStatus() *CmdStatus {
	cfg := entity.Config{}.FromFile()
	var user_repo repository.UserRepository = nil
	var auth_uc auth_usecase.UseCase = nil
	var userCli pbUser.UserApiClient = nil

	var server_repo repository.ServerRepository = nil
	var server_uc server_usecase.UseCase = nil
	var serverCli pbServer.ServerApiClient = nil

	if cfg != nil {
		var err error
		userCli, err = pbUser.NewUserApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server:", err)
		}

		serverCli, err = pbServer.NewServerApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		user_repo = user_repository.New(userCli)
		auth_uc = auth_usecase.New(user_repo)

		server_repo = server_repository.New(serverCli)
		server_uc = server_usecase.New(server_repo)
	}

	c := &CmdStatus{
		auth_uc:   auth_uc,
		server_uc: server_uc,
	}
	c.Command = &cobra.Command{
		Use:   "status",
		Short: "Get status",
		Long:  "Get status",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdStatus) runCommand(cmd *cobra.Command, args []string) error {
	if c.auth_uc == nil || c.server_uc == nil {
		return errors.New("You haven't configure dply-server host. Run `dply config edit`")
	}
	cfg := entity.Config{}.FromFile()
	if cfg == nil {
		fmt.Println("dply-server: Not Connect")
		fmt.Println("login status: Not Login")
		return nil
	}
	isConnectServer := c.server_uc.Status()
	isLogin, userInfo := c.auth_uc.GetStatus()
	if isConnectServer {
		fmt.Println("dply-server: Connected (" + cfg.DplyServerHost + ")")
	} else {
		fmt.Println("dply-server: Not Connect")
	}

	if isLogin {
		fmt.Println("login status: Login (" + userInfo.Email + ")")
	} else {
		fmt.Println("login status: Not Login")
	}

	return nil
}

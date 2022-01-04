package cli_envar

import (
	"log"

	"github.com/herryg91/dply/dply/app/repository"
	deploy_usecase "github.com/herryg91/dply/dply/app/usecase/deploy"
	pbDeploy "github.com/herryg91/dply/dply/clients/grst/deploy"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/deploy_repository"
	"github.com/spf13/cobra"
)

type CmdDeploy struct {
	*cobra.Command
}

func New() *CmdDeploy {
	cfg := entity.Config{}.FromFile()
	var deploy_repo repository.DeployRepository = nil
	var deploy_uc deploy_usecase.UseCase = nil
	var deployCli pbDeploy.DeployApiClient = nil
	if cfg != nil {
		var err error
		deployCli, err = pbDeploy.NewDeployApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		deploy_repo = deploy_repository.New(deployCli)
		deploy_uc = deploy_usecase.New(deploy_repo)
	}

	c := &CmdDeploy{}
	c.Command = &cobra.Command{
		Use:   "deploy",
		Short: "Manage deployment",
		Long:  "Manage deployment",
	}

	c.AddCommand(newDeployImage(deploy_uc).Command)
	c.AddCommand(newDeployReload(deploy_uc).Command)
	return c
}

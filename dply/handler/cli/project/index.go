package cli_project

import (
	"log"

	"github.com/herryg91/dply/dply/app/repository"
	project_usecase "github.com/herryg91/dply/dply/app/usecase/project"
	pbProject "github.com/herryg91/dply/dply/clients/grst/project"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/project_repository"
	"github.com/spf13/cobra"
)

type CmdConfig struct {
	*cobra.Command
}

func New() *CmdConfig {
	c := &CmdConfig{}
	c.Command = &cobra.Command{
		Use:   "project",
		Short: "Configure project",
		Long:  "Configure project",
	}
	cfg := entity.Config{}.FromFile()
	var project_repo repository.ProjectRepository = nil
	var project_uc project_usecase.UseCase = nil
	var projectCli pbProject.ProjectApiClient = nil
	if cfg != nil {
		var err error
		projectCli, err = pbProject.NewProjectApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		project_repo = project_repository.New(projectCli)
		project_uc = project_usecase.New(project_repo)
	}

	c.AddCommand(newCmdProjectChange().Command)
	c.AddCommand(newCmdProjectList(project_uc).Command)
	c.AddCommand(newCmdProjectCreate(project_uc).Command)
	c.AddCommand(newCmdProjectDelete(project_uc).Command)
	return c
}

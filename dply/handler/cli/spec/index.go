package cli_spec

import (
	"log"

	"github.com/herryg91/dply/dply/app/repository"
	affinity_usecase "github.com/herryg91/dply/dply/app/usecase/affinity"
	deployment_config_usecase "github.com/herryg91/dply/dply/app/usecase/deployment-config"
	envar_usecase "github.com/herryg91/dply/dply/app/usecase/envar"
	port_usecase "github.com/herryg91/dply/dply/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply/app/usecase/scale"
	pbSpec "github.com/herryg91/dply/dply/clients/grst/spec"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/repository/spec_repository"
	"github.com/spf13/cobra"
)

type CmdSpec struct {
	*cobra.Command
}

func New() *CmdSpec {
	cfg := entity.Config{}.FromFile()
	var spec_repo repository.SpecRepository = nil
	var specCli pbSpec.SpecApiClient = nil

	var scale_uc scale_usecase.UseCase = nil
	var envar_uc envar_usecase.UseCase = nil
	var port_uc port_usecase.UseCase = nil
	var affinity_uc affinity_usecase.UseCase = nil
	var deployment_config_uc deployment_config_usecase.UseCase = nil

	if cfg != nil {
		var err error
		specCli, err = pbSpec.NewSpecApiGrstClient(cfg.DplyServerHost, nil)
		if err != nil {
			log.Panicln("Failed to initialized cli for dply-server", err)
		}

		spec_repo = spec_repository.New(specCli)
		scale_uc = scale_usecase.New(spec_repo)
		envar_uc = envar_usecase.New(spec_repo)
		port_uc = port_usecase.New(spec_repo)
		affinity_uc = affinity_usecase.New(spec_repo)
		deployment_config_uc = deployment_config_usecase.New(spec_repo)
	}

	c := &CmdSpec{}
	c.Command = &cobra.Command{
		Use:   "spec",
		Short: "Manage deployment spec: envar, port, etc.",
		Long:  "Manage deployment spec: envar, port, etc.",
	}

	c.AddCommand(newSpecGet(cfg, envar_uc, port_uc, scale_uc, affinity_uc).Command)
	c.AddCommand(newSpecScalingGet(cfg, scale_uc).Command)
	c.AddCommand(newSpecScalingEdit(cfg, scale_uc).Command)
	c.AddCommand(newSpecEnvarGet(cfg, envar_uc).Command)
	c.AddCommand(newSpecEnvarEdit(cfg, envar_uc).Command)
	c.AddCommand(newSpecEnvarSet(cfg, envar_uc).Command)
	c.AddCommand(newSpecPortGet(cfg, port_uc).Command)
	c.AddCommand(newSpecPortEdit(cfg, port_uc).Command)
	c.AddCommand(newSpecAffinityGet(cfg, affinity_uc).Command)
	c.AddCommand(newSpecAffinityEdit(cfg, affinity_uc).Command)
	c.AddCommand(newSpecDeploymentConfigGet(cfg, deployment_config_uc).Command)
	c.AddCommand(newSpecDeploymentConfigEdit(cfg, deployment_config_uc).Command)

	return c
}

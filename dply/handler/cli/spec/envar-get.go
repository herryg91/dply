package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	envar_usecase "github.com/herryg91/dply/dply/app/usecase/envar"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecEnvarGet struct {
	*cobra.Command
	envar_uc envar_usecase.UseCase

	project string
	env     string
	name    string
}

func newSpecEnvarGet(cfg *entity.Config, envar_uc envar_usecase.UseCase) *CmdSpecEnvarGet {
	c := &CmdSpecEnvarGet{project: cfg.Project, envar_uc: envar_uc}
	c.Command = &cobra.Command{
		Use:     "envar-get",
		Aliases: []string{"eg"},
		Short:   "View environment variables",
		Long:    "View environment variables",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecEnvarGet) runCommand(cmd *cobra.Command, args []string) error {
	service_yaml_data, err := serviceYaml.GetServiceYAML("service.yaml")
	if err == nil {
		if service_yaml_data.Project != "" {
			c.project = service_yaml_data.Project
		}
	}

	if c.envar_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.env == "" {
		return errors.New("`--env / -e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	resp, err := c.envar_uc.Get(c.project, c.env, c.name)
	if err != nil {
		if errors.Is(err, envar_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	jsonData, _ := json.MarshalIndent(resp.Variables, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

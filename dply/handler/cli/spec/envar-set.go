package cli_spec

import (
	"errors"
	"fmt"
	"io/ioutil"

	envar_usecase "github.com/herryg91/dply/dply/app/usecase/envar"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecEnvarSet struct {
	*cobra.Command
	envar_uc envar_usecase.UseCase

	project string
	env     string
	name    string
	key     string
	value   string
	file    string
}

func newSpecEnvarSet(cfg *entity.Config, envar_uc envar_usecase.UseCase) *CmdSpecEnvarSet {
	c := &CmdSpecEnvarSet{
		envar_uc: envar_uc,
		project:  cfg.Project,
	}
	c.Command = &cobra.Command{
		Use:     "envar-set",
		Aliases: []string{"es"},
		Short:   "Set environment variables value",
		Long:    "Set environment variables value",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	c.Command.Flags().StringVarP(&c.key, "key", "k", "", "key")
	c.Command.Flags().StringVarP(&c.value, "value", "v", "", "value")
	c.Command.Flags().StringVarP(&c.file, "file", "f", "", "value from file")
	return c
}

func (c *CmdSpecEnvarSet) runCommand(cmd *cobra.Command, args []string) error {
	service_yaml_data, err := serviceYaml.GetServiceYAML("service.yaml")
	if err == nil {
		if service_yaml_data.Project != "" {
			c.project = service_yaml_data.Project
		}
	}

	if c.envar_uc == nil {
		return errors.New("You haven't configure config. command: `dply-cli config --server=<dply_server_host>`")
	} else if c.env == "" {
		return errors.New("`-e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	} else if c.key == "" {
		return errors.New("`--key / -k` is required")
	} else if c.file == "" && c.value == "" {
		return errors.New("`--value or --file` is required")
	}

	current_envar, err := c.envar_uc.Get(c.project, c.env, c.name)
	if err != nil {
		if errors.Is(err, envar_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	if c.value != "" {
		current_envar.Variables[c.key] = c.value
	} else if c.file != "" {
		from_file_content, err := ioutil.ReadFile(c.file)
		if err != nil {
			return err
		}
		current_envar.Variables[c.key] = string(from_file_content)
	} else {
		fmt.Println("Nothing to change")
		return nil
	}

	fmt.Println("environment variable " + c.name + " (" + c.env + "): '" + c.key + "' succesfully updated")
	c.envar_uc.Upsert(*current_envar)

	return nil
}

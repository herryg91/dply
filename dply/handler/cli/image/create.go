package cli_image

import (
	"errors"
	"strings"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"

	"github.com/spf13/cobra"
)

type CmdImageCreate struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	project         string
	name            string
	description     string
	registry_prefix string
	build_args      []string
}

func newCmdImageCreate(cfg *entity.Config, image_uc image_usecase.UseCase) *CmdImageCreate {
	c := &CmdImageCreate{project: cfg.Project, image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "create",
		Short: "Create image",
		Long:  "Create image",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "service (repository) name of image")
	c.Command.Flags().StringVarP(&c.description, "desc", "d", "", "image description")
	c.Command.Flags().StringVarP(&c.registry_prefix, "prefix", "p", "", "registry prefix")
	c.Command.Flags().StringSliceVarP(&c.build_args, "arg", "a", []string{}, "build arguments")

	return c
}

func (c *CmdImageCreate) runCommand(cmd *cobra.Command, args []string) error {
	service_yaml_data, err := serviceYaml.GetServiceYAML("service.yaml")
	if err == nil {
		if service_yaml_data.Project != "" {
			c.project = service_yaml_data.Project
		}
	}

	if c.image_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.description == "" {
		return errors.New("`--desc / -d` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}
	cfg := entity.Config{}.FromFile()
	svcYaml, _ := serviceYaml.GetServiceYAML("service.yaml")
	// Olah registry prefix
	if c.registry_prefix == "" && svcYaml.Category != "" {
		c.registry_prefix = svcYaml.Category
	}
	if c.registry_prefix != "" {
		cfg.RegistryTagPrefix += "/" + c.registry_prefix
	}

	// Olah build_args
	build_args_formatted := map[string]*string{}
	for _, a := range c.build_args {
		splitted_args := strings.SplitN(a, "=", 2)
		if len(splitted_args) != 2 {
			continue
		}
		val := splitted_args[1]
		if len(val) >= 2 {
			if string(val[0]) == `"` && string(val[len(val)-1]) == `"` {
				val = val[1 : len(val)-1]
			}
		}
		build_args_formatted[splitted_args[0]] = &val
	}
	for _, v := range svcYaml.BuildArgs {
		build_args_formatted[v.Name] = &v.Value
	}

	err = c.image_uc.Create(c.project, c.name, cfg.RegistryTagPrefix, c.description, build_args_formatted)
	if err != nil {
		return err
	}

	return nil
}

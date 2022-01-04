package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	port_usecase "github.com/herryg91/dply/dply/app/usecase/port"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecPortGet struct {
	*cobra.Command
	port_uc port_usecase.UseCase

	env  string
	name string
}

func newSpecPortGet(port_uc port_usecase.UseCase) *CmdSpecPortGet {
	c := &CmdSpecPortGet{port_uc: port_uc}
	c.Command = &cobra.Command{
		Use:     "port-get",
		Aliases: []string{"pg"},
		Short:   "View port specification",
		Long:    "View port specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecPortGet) runCommand(cmd *cobra.Command, args []string) error {
	if c.port_uc == nil {
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

	resp, err := c.port_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, port_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	jsonData, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

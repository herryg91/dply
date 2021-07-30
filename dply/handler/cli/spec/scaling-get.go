package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	scale_usecase "github.com/herryg91/dply/dply/app/usecase/scale"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecScalingGet struct {
	*cobra.Command
	scale_uc scale_usecase.UseCase

	env  string
	name string
}

func newSpecScalingGet(scale_uc scale_usecase.UseCase) *CmdSpecScalingGet {
	c := &CmdSpecScalingGet{scale_uc: scale_uc}
	c.Command = &cobra.Command{
		Use:     "scaling-get",
		Aliases: []string{"sg"},
		Short:   "View scaling specification",
		Long:    "View scaling specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecScalingGet) runCommand(cmd *cobra.Command, args []string) error {
	if c.scale_uc == nil {
		return errors.New("You haven't configure setting. command: `dply-cli setting --server=<dply_server_host>`")
	} else if c.env == "" {
		return errors.New("`--env / -e` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	resp, err := c.scale_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, scale_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	jsonData, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

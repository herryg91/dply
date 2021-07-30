package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	affinity_usecase "github.com/herryg91/dply/dply/app/usecase/affinity"
	envar_usecase "github.com/herryg91/dply/dply/app/usecase/envar"
	port_usecase "github.com/herryg91/dply/dply/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply/app/usecase/scale"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecGet struct {
	*cobra.Command
	envar_uc    envar_usecase.UseCase
	port_uc     port_usecase.UseCase
	scale_uc    scale_usecase.UseCase
	affinity_uc affinity_usecase.UseCase

	env  string
	name string
}

func newSpecGet(envar_uc envar_usecase.UseCase, port_uc port_usecase.UseCase, scale_uc scale_usecase.UseCase, affinity_uc affinity_usecase.UseCase) *CmdSpecGet {
	c := &CmdSpecGet{
		envar_uc:    envar_uc,
		port_uc:     port_uc,
		scale_uc:    scale_uc,
		affinity_uc: affinity_uc,
	}
	c.Command = &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "View all specification",
		Long:    "View all specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecGet) runCommand(cmd *cobra.Command, args []string) error {
	if c.envar_uc == nil || c.port_uc == nil {
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
	respEnvar, err := c.envar_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, envar_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	respPort, err := c.port_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, port_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	respScale, err := c.scale_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, scale_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	respAffinity, err := c.affinity_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, affinity_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	resp := map[string]interface{}{
		"variables": respEnvar.Variables,
		"scaling":   respScale,
		"ports":     respPort.Ports,
		"affinity":  respAffinity,
	}

	jsonData, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

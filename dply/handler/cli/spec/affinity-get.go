package cli_spec

import (
	"encoding/json"
	"errors"
	"fmt"

	affinity_usecase "github.com/herryg91/dply/dply/app/usecase/affinity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdSpecAffinityGet struct {
	*cobra.Command
	affinity_uc affinity_usecase.UseCase

	env  string
	name string
}

func newSpecAffinityGet(affinity_uc affinity_usecase.UseCase) *CmdSpecAffinityGet {
	c := &CmdSpecAffinityGet{affinity_uc: affinity_uc}
	c.Command = &cobra.Command{
		Use:     "affinity-get",
		Aliases: []string{"ag"},
		Short:   "View affinity specification",
		Long:    "View affinity specification",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.env, "env", "e", "", "environment/namespace")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "deployment name")
	return c
}

func (c *CmdSpecAffinityGet) runCommand(cmd *cobra.Command, args []string) error {
	if c.affinity_uc == nil {
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

	resp, err := c.affinity_uc.Get(c.env, c.name)
	if err != nil {
		if errors.Is(err, affinity_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}

	jsonData, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(jsonData))
	return nil
}

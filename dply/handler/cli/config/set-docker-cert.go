package cli_config

import (
	"fmt"

	"github.com/herryg91/dply/dply/entity"
	"github.com/spf13/cobra"
)

type CmdConfigSetDockerCert struct {
	*cobra.Command
}

func newConfigSetDockerCert() *CmdConfigSetDockerCert {
	c := &CmdConfigSetDockerCert{}
	c.Command = &cobra.Command{
		Use:   "set-docker-cert",
		Short: "Set docker engine certificates folder location (ca.pem, cert.pem, key.pem)",
		Long:  "Set docker engine certificates folder location (ca.pem, cert.pem, key.pem)",
	}
	c.RunE = c.runCommand
	return c
}

func (c *CmdConfigSetDockerCert) runCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require value parameter")
	}

	new_docker_cert_loc := args[0]
	cfg := entity.Config{}.FromFile()
	old_docker_cert_loc := cfg.DockerCertificatesPath
	if cfg.DockerCertificatesPath != new_docker_cert_loc {
		cfg.DockerCertificatesPath = new_docker_cert_loc
		err := cfg.SaveConfig()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Docker engine's certificates folder was changed, %s -> %s", old_docker_cert_loc, new_docker_cert_loc))
	} else {
		fmt.Println("Nothing to update")
	}

	return nil
}

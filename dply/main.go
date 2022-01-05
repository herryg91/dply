package main

import (
	"os"

	cli "github.com/herryg91/dply/dply/handler/cli"
	cli_config "github.com/herryg91/dply/dply/handler/cli/config"
	cli_deploy "github.com/herryg91/dply/dply/handler/cli/deploy"
	cli_image "github.com/herryg91/dply/dply/handler/cli/image"
	cli_project "github.com/herryg91/dply/dply/handler/cli/project"
	cli_spec "github.com/herryg91/dply/dply/handler/cli/spec"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{Use: "dply", Short: "dply", Long: "dply"}
	rootCmd.AddCommand(cli.NewCmdStatus().Command)
	rootCmd.AddCommand(cli.NewCmdLogin().Command)
	rootCmd.AddCommand(cli.NewCmdLogout().Command)
	rootCmd.AddCommand(cli_config.New().Command)
	rootCmd.AddCommand(cli_project.New().Command)

	rootCmd.AddCommand(cli_image.New().Command)
	rootCmd.AddCommand(cli_spec.New().Command)
	rootCmd.AddCommand(cli_deploy.New().Command)

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}

}

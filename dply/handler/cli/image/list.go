package cli_image

import (
	"errors"
	"fmt"
	"os"
	"time"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CmdImageList struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	project string
	name    string
	page    int
	size    int
}

func newCmdImageList(cfg *entity.Config, image_uc image_usecase.UseCase) *CmdImageList {
	c := &CmdImageList{project: cfg.Project, image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "list",
		Short: "Get container images",
		Long:  "Get container images",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "repository name (default: parameter <name> in service.yaml file)")
	c.Command.Flags().IntVar(&c.page, "page", 1, "page")
	c.Command.Flags().IntVar(&c.size, "size", 10, "size")
	return c
}

func (c *CmdImageList) runCommand(cmd *cobra.Command, args []string) error {
	if c.image_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("parameter --name or -n is required")
		}
		c.name = data.Name
	}

	datas, err := c.image_uc.GetList(c.project, c.name, c.page, c.size)
	if err != nil {
		if errors.Is(err, image_usecase.ErrUnauthorized) {
			logrus.Errorln(err.Error())
			return nil
		}
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Digest",
		"Description",
		"Created At",
		"Notes",
		// "Full Image",
	})

	for _, data := range datas {
		tableData := []string{
			data.Digest,
			data.Description,
			data.CreatedAt.Format(time.RFC3339),
			data.Notes,
			// data.Image,
		}

		table.Append(tableData)
	}
	fmt.Println("repository: " + c.name)
	table.Render()

	return nil
}

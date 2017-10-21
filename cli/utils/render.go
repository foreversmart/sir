package utils

import (
	"sir/models"

	"github.com/yanhao1991/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

func RenderTaskConfig(taskConfig *models.TaskConfig, c *cli.Context) {
	table := tablewriter.NewWriter(c.App.Writer)

	table.SetRowLine(false)

	table.AppendBulk([][]string{
		[]string{Style.Title("NAME"), Style.Bold(taskConfig.Name)},
		[]string{Style.Title("EXEC"), taskConfig.Cmd},
		[]string{Style.Title("WATCH"), Format.Enabled(taskConfig.Watch)},
		[]string{Style.Title("WATCH_DIR"), taskConfig.WatchDir},
		// []string{Style.Title("ENV"), taskConfig.Env},
	})

	table.Render()
}

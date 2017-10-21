package utils

import (
	"sir/models"
	"strconv"

	"github.com/yanhao1991/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

func RenderTaskConfig(taskConfig *models.TaskConfig, c *cli.Context) {
	table := tablewriter.NewWriter(c.App.Writer)

	table.SetRowLine(false)

	table.Append([]string{Style.Title("NAME"), Style.Bold(taskConfig.Name)})
	table.Append([]string{Style.Title("EXEC"), taskConfig.Cmd})
	table.Append([]string{Style.Title("WATCH"), Format.Enabled(taskConfig.Watch)})
	table.Append([]string{Style.Title("WATCH_DIR"), taskConfig.WatchDir})
	table.Append([]string{Style.Title("WORKSPACE"), taskConfig.Workspace})
	table.Append([]string{Style.Title("USER"), taskConfig.User})
	table.Append([]string{Style.Title("GROUP"), taskConfig.Group})
	table.Append([]string{Style.Title("PRIORITY"), strconv.Itoa(taskConfig.Priority)})
	table.Append([]string{Style.Title("AUTO_START"), Format.Enabled(taskConfig.AutoStart)})
	table.Append([]string{Style.Title("AUTO_RESTART"), Format.Enabled(taskConfig.AutoRestart)})
	table.Append([]string{Style.Title("RESTART_INTERVAL"), strconv.Itoa(taskConfig.RestartInterval)})
	table.Append([]string{Style.Title("RESTART_COUNT"), string(taskConfig.RestartCount)})

	// ENV
	envs := Format.KVMap(taskConfig.Env)
	for i, env := range envs {
		key := " "
		if i == 0 {
			key = Style.Title("ENV")
		}
		table.Append([]string{key, env})
	}

	// RULES
	for i, rule := range taskConfig.Rules {
		key := " "
		if i == 0 {
			key = Style.Title("RULES")
		}
		table.Append([]string{key, Format.KV(rule.Type, strconv.FormatFloat(rule.Threshold, 'f', 2, 64))})
	}

	table.Append([]string{Style.Title("ERR_LOG"), taskConfig.LogConfigs.ErrLogPath})
	table.Append([]string{Style.Title("STD_LOG"), taskConfig.LogConfigs.StdLogPath})
	table.Append([]string{Style.Title("LOG_ROTATE"), taskConfig.LogConfigs.RotateType})
	table.Append([]string{Style.Title("LOG_ROTATE_LIMIT"), strconv.Itoa(taskConfig.LogConfigs.Limit)})

	table.Append([]string{Style.Title("CREATED_AT"), taskConfig.CTime.String()})

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

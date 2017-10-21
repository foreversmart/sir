package utils

import (
	"fmt"
	"sir/models"
	"sort"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/yanhao1991/tablewriter"
	cli "gopkg.in/urfave/cli.v1"
)

func RenderTaskConfig(taskConfig *models.TaskConfig, c *cli.Context) {
	fmt.Println(Style.Header(" # TASK CONFIG INFO:"))

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

	table.Append([]string{Style.Title("CREATED_AT"), taskConfig.CTime.Format("2006-01-02T15:04:05Z07:00")})

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

type TaskSlice []models.Task

func (s TaskSlice) Len() int      { return len(s) }
func (s TaskSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s TaskSlice) Less(i, j int) bool {
	vi := 0
	vj := 0
	if s[i].TaskState == nil {
		vi = -1
	}
	if s[j].TaskState == nil {
		vj = -1
	}
	return vi > vj
}

func RenderTaskList(list []models.Task, c *cli.Context) {
	println(len(list), "?")

	sort.Sort(TaskSlice(list))

	table := tablewriter.NewWriter(c.App.Writer)
	table.SetHeader([]string{"task name", "priority", "pid", "status", "restart", "uptime", "cpu", "mem", "user", "watching"})
	table.SetHeaderColor(
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
		tablewriter.Color(tablewriter.FgCyanColor),
	)

	for _, t := range list {
		v := make([]string, 10)

		v[0] = Style.Bold(t.Name)
		v[1] = strconv.Itoa(t.Priority)
		v[2] = "-"
		v[3] = Style.Fail("stoped")
		v[4] = strconv.Itoa(t.RestartCount)
		v[5] = "-"
		v[6] = "-"
		v[7] = "-"
		v[8] = t.User
		v[9] = Format.Enabled(t.Watch)

		if t.TaskState != nil {
			v[0] = Style.Title(t.Name)
			v[2] = strconv.Itoa(int(t.Pid))
			v[3] = Style.Success("running")

			seconds := (int64(time.Now().Sub(time.Unix(t.UpTime, 0))) / 1e9) * 1e9
			v[5] = time.Duration(seconds).String()
			v[6] = humanize.FormatFloat("###.##", t.CpuPercent) + " %"
			v[7] = humanize.Bytes(t.Mem)
		}

		table.Append(v)
	}

	table.Render() // Send output
}

func RenderTask(task *models.Task, c *cli.Context) {
	println(task.Name)
}

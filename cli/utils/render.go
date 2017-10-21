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
	fmt.Println(Style.Header(" # TASK CONFIG INFO "))

	table := tablewriter.NewWriter(c.App.Writer)

	table.SetRowLine(false)

	table.Append([]string{Style.Title("NAME"), Style.Title(taskConfig.Name)})
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
	table.Append([]string{Style.Title("RESTART_COUNT"), strconv.Itoa(taskConfig.RestartCount)})

	// ENV
	for i, env := range taskConfig.Env {
		key := " "
		if i == 0 {
			key = Style.Title("ENV")
		}
		table.Append([]string{key, env})
	}

	if len(taskConfig.Env) == 0 {
		table.Append([]string{Style.Title("ENV"), ""})
	}

	// RULES
	for i, rule := range taskConfig.Rules {
		key := " "
		if i == 0 {
			key = Style.Title("RULES")
		}
		table.Append([]string{key, Format.KV(rule.Type, strconv.FormatFloat(rule.Threshold, 'f', 2, 64))})
	}
	if len(taskConfig.Rules) == 0 {
		table.Append([]string{Style.Title("RULES"), ""})
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

func RenderTaskState(task *models.TaskState, c *cli.Context) {
	fmt.Println(Style.Header(" # TASK PROCESS INFO "))

	table := tablewriter.NewWriter(c.App.Writer)

	table.SetRowLine(false)

	table.Append([]string{Style.Title("PID"), Style.Title(strconv.Itoa(int(task.Pid)))})
	table.Append([]string{Style.Title("CPU"), humanize.FormatFloat("###.##", task.CpuPercent) + " %"})
	table.Append([]string{Style.Title("MEMORY"), humanize.Bytes(task.Mem)})
	table.Append([]string{Style.Title("LOAD"), humanize.FormatFloat("###.##", task.Load)})
	table.Append([]string{Style.Title("STAT"), task.Stat})

	table.Append([]string{Style.Title("DISK_IO"), "R    " + humanize.Bytes(task.IoCounter.ReadBytes)})
	table.Append([]string{"", "W    " + humanize.Bytes(task.IoCounter.WriteBytes)})

	table.Append([]string{Style.Title("NETWORK_IO"), "SENT " + humanize.Bytes(task.Net.BytesSent)})
	table.Append([]string{"", "RECV " + humanize.Bytes(task.Net.BytesSent)})

	seconds := (int64(time.Now().Sub(time.Unix(task.UpTime, 0))) / 1e9) * 1e9
	table.Append([]string{Style.Title("UP_TIME"), time.Duration(seconds).String()})

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()

}

func RenderTask(task *models.Task, c *cli.Context) {
	println()
	if task.TaskState != nil {
		fmt.Println(Style.Success("[INFO]"), Style.Title(task.Name), "TASK PROCESS IS", Style.Success("RUNNING"), "\n")
		RenderTaskState(task.TaskState, c)
		println()
	} else {
		fmt.Println(Style.Success("[INFO]"), Style.Title(task.Name), "TASK PROCESS IS", Style.Fail("STOPED"), "\n")
	}

	RenderTaskConfig(task.TaskConfig, c)
}

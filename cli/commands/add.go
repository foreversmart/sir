package commands

import (
	"fmt"
	"net/http"
	"os"
	"sir/cli/config"
	"sir/cli/opts"
	"sir/cli/utils"
	"sir/lib/httpclient"
	"sir/models"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdAdd = cli.Command{
	Name:      "add",
	UsageText: "<cmd> --name <taskname>",
	Category:  string(ConfigCategory),
	Usage:     "create a task config by params",
	Action:    ActionAdd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "must have, assign name for new task config",
		},
		cli.StringFlag{
			Name:  "watch",
			Usage: "enable watch & reload function, value is the watch folder path, default is false",
		},
		cli.StringFlag{
			Name:  "env",
			Usage: "run time env., format: \"key=value&key=value\"",
		},
		cli.StringFlag{
			Name:  "workspace",
			Usage: "enable workspace mode, value is workspace folder path, default is false",
		},
		cli.StringFlag{
			Name:  "user",
			Usage: "run as which user, default is current user",
		},
		cli.StringFlag{
			Name:  "group",
			Usage: "run as which group, default is current group",
		},
		cli.IntFlag{
			Name:  "priority",
			Usage: "task priority, decide the auto start seq., default is 0",
		},
		cli.BoolFlag{
			Name:  "autostart",
			Usage: "enable auto start when daemon start, default is false",
		},
		cli.BoolFlag{
			Name:  "autorestart",
			Usage: "enable auto restart when crash, default is false",
		},
		cli.IntFlag{
			Name:  "restartinterval",
			Usage: "enable auto restart when crash, default is 0s",
		},
		cli.StringFlag{
			Name:  "errlog",
			Usage: "assign err log folder path, default is ~/.sir/logs/",
		},
		cli.StringFlag{
			Name:  "stdlog",
			Usage: "assign std log folder path, default is ~/.sir/logs/",
		},
		cli.StringFlag{
			Name:  "logrotate",
			Usage: "log file rotate type: day|size, default is day",
		},
		cli.Float64Flag{
			Name:  "loglimit",
			Usage: "log file rotate limit, unit is DAY or MB",
		},
		cli.StringFlag{
			Name:  "rules",
			Usage: "set auto restart rules, format is <TYPE>=<THRESHOLD>&<TYPE>=<THRESHOLD>, types: cpu|mem|load|uptime, threshold units:percentage|MB|.|seconds",
		},
	},
}

func ActionAdd(c *cli.Context) error {

	// 获取参数
	opts := opts.ParseAddOpts(c)
	if !opts.IsValid() {
		fmt.Printf("ERROR: params format error, please check \"sir add -h\"\n")
		os.Exit(0)
	}

	// 获取可执行文件路径
	cmd, err := utils.ExecFileAbsPath(opts.Cmd)
	if err != nil {
		fmt.Printf("ERROR: can not find exec bin, %v \n", err)
		os.Exit(0)
	}
	opts.Cmd = cmd

	// 调用api
	var response map[string]models.TaskConfig
	httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/add"), opts, response)
	// TODO handle error

	taskConfig := response["data"]

	// 输出结果
	utils.RenderTaskConfig(&taskConfig, c)

	return nil
}

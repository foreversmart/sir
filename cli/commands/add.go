package commands

import (
	"fmt"
	"os"
	"sir/cli/mock"
	"sir/cli/opts"
	"sir/cli/utils"

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
			Usage: "assign name for new task config",
		},
	},
}

func ActionAdd(c *cli.Context) error {

	// 获取参数
	opts := opts.ParseAddOpts(c)
	if !opts.IsValid() {
		fmt.Printf("ERROR: Missing params 'sir add <cmd> --name <taskname>'\n")
		os.Exit(0)
	}

	// 获取可执行文件路径
	cmd, err := utils.ExecFileAbsPath(opts.Cmd)
	if err != nil {
		fmt.Printf("ERROR: can not find exec bin, %v \n", err)
		os.Exit(0)
	}
	opts.Cmd = cmd

	// TODO send command to sird daemon
	taskConfig := mock.GetTaskConfig()

	// 输出结果
	utils.RenderTaskConfig(&taskConfig, c)

	return nil
}

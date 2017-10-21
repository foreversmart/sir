package commands

import (
	"sir/cli/mock"
	"sir/cli/utils"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdShow = cli.Command{
	Name:      "show",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "describe all parameters of a task",
	Action:    ActionShow,
}

func ActionShow(c *cli.Context) error {
	// TODO call sird api

	task := mock.GetTask()

	utils.RenderTask(&task, c)

	return nil
}

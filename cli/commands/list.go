package commands

import (
	"sir/cli/mock"
	"sir/cli/utils"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdList = cli.Command{
	Name:      "list",
	UsageText: "",
	Category:  string(TaskCategory),
	Usage:     "list all tasks",
	Action:    ActionList,
}

func ActionList(c *cli.Context) error {
	// TODO call sir daemon
	list := mock.GetTaskList()

	utils.RenderTaskList(list, c)

	return nil
}

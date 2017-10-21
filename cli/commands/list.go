package commands

import (
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
	return nil
}

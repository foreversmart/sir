package commands

import (
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
	return nil
}

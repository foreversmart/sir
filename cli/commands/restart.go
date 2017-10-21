package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdRestart = cli.Command{
	Name:      "restart",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "restart a task",
	Action:    ActionRestart,
}

func ActionRestart(c *cli.Context) error {
	return nil
}

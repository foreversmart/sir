package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStart = cli.Command{
	Name:      "start",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "start and daemonize a task",
	Action:    ActionStart,
}

func ActionStart(c *cli.Context) error {
	return nil
}

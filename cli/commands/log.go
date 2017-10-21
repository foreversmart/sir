package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdLog = cli.Command{
	Name:      "log",
	UsageText: "[<task>] [-from <fromtime>] [-to <endtime>]",
	Category:  string(TaskCategory),
	Usage:     "stream logs file. Default stream all tasks logs",
	Action:    ActionLog,
}

func ActionLog(c *cli.Context) error {
	return nil
}

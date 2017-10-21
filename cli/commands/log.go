package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdLog = cli.Command{
	Name:      "log",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionLog,
}

func ActionLog(c *cli.Context) error {
	return nil
}

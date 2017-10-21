package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdRestart = cli.Command{
	Name:      "restart",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionRestart,
}

func ActionRestart(c *cli.Context) error {
	return nil
}

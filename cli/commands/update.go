package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdUpdate = cli.Command{
	Name:      "update",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionUpdate,
}

func ActionUpdate(c *cli.Context) error {
	return nil
}

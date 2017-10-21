package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdRemove = cli.Command{
	Name:      "remove",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionRemove,
}

func ActionRemove(c *cli.Context) error {
	return nil
}

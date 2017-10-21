package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdShow = cli.Command{
	Name:      "show",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionShow,
}

func ActionShow(c *cli.Context) error {
	return nil
}

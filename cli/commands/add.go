package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdAdd = cli.Command{
	Name:      "add",
	UsageText: "<cmd> [extra params]",
	Category:  string(ConfigCategory),
	Usage:     "create a task config by params",
	Action:    ActionAdd,
}

func ActionAdd(c *cli.Context) error {
	return nil
}

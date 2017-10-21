package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdUpdate = cli.Command{
	Name:      "update",
	UsageText: "<task>",
	Category:  string(ConfigCategory),
	Usage:     "update task configs with default editor",
	Action:    ActionUpdate,
}

func ActionUpdate(c *cli.Context) error {
	return nil
}

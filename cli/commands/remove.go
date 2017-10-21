package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdRemove = cli.Command{
	Name:      "remove",
	UsageText: "<task>",
	Category:  string(ConfigCategory),
	Usage:     "remove a task config by name",
	Action:    ActionRemove,
}

func ActionRemove(c *cli.Context) error {
	// TODO
	return nil
}

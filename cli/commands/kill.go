package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdKill = cli.Command{
	Name:      "kill",
	UsageText: "",
	Category:  string(ServiceCategory),
	Usage:     "kill sir daemon process",
	Action:    ActionKill,
}

func ActionKill(c *cli.Context) error {
	return nil
}

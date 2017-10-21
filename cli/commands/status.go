package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStatus = cli.Command{
	Name:      "status",
	UsageText: "",
	Category:  string(ServiceCategory),
	Usage:     "show sir daemon detail status",
	Action:    ActionStatus,
}

func ActionStatus(c *cli.Context) error {
	return nil
}

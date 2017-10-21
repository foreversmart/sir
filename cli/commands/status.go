package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStatus = cli.Command{
  Name: "status",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionStatus,
}

func ActionStatus(c *cli.Context) error {
	return nil
}

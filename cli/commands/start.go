package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStart = cli.Command{
  Name: "start",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionStart,
}

func ActionStart(c *cli.Context) error {
	return nil
}

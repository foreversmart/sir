package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStop = cli.Command{
  Name: "stop",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionStop,
}

func ActionStop(c *cli.Context) error {
	return nil
}

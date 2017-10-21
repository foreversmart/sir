package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdKill = cli.Command{
  Name: "kill",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionKill,
}

func ActionKill(c *cli.Context) error {
	return nil
}

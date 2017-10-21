package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdSend = cli.Command{
  Name: "send",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionSend,
}

func ActionSend(c *cli.Context) error {
	return nil
}

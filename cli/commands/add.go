package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdAdd = cli.Command{
  Name: "add",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: ActionAdd,
}

func ActionAdd(c *cli.Context) error {
	return nil
}

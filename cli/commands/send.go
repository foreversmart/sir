package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdSend = cli.Command{
	Name:      "send",
	UsageText: "<task> -m \"xxx\"",
	Category:  string(TaskCategory),
	Usage:     "send stdin to a task process",
	Action:    ActionSend,
}

func ActionSend(c *cli.Context) error {
	return nil
}

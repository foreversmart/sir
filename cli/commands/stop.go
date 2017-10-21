package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStop = cli.Command{
	Name:      "stop",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "stop a task (to start it again, do sir restart <task>)",
	Action:    ActionStop,
}

func ActionStop(c *cli.Context) error {
	return nil
}

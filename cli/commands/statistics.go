package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStatistics = cli.Command{
	Name:      "statistics",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionStatistics,
}

func ActionStatistics(c *cli.Context) error {
	return nil
}

package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdStatistics = cli.Command{
	Name:      "statistics",
	UsageText: "[<task>] [-from <fromtime>] [-to <endtime>]",
	Category:  string(ServiceCategory),
	Usage:     "show statistics, default sir daemon statistics",
	Action:    ActionStatistics,
}

func ActionStatistics(c *cli.Context) error {
	return nil
}

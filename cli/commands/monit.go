package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdMonit = cli.Command{
	Name:      "monit",
	UsageText: "",
	Category:  "",
	Usage:     "",
	Action:    ActionMonit,
}

func ActionMonit(c *cli.Context) error {
	return nil
}

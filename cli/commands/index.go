package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

type CommandCategory string

const (
	ConfigCategory  CommandCategory = "Config Management"
	TaskCategory    CommandCategory = "Task Management"
	ServiceCategory CommandCategory = "Service Management"
)

var CliCmds = []cli.Command{
	// config cmds
	CmdAdd,
	CmdRemove,
	CmdUpdate,

	// task cmds
	CmdStart,
	CmdStop,
	CmdRestart,
	CmdList,
	CmdShow,
	CmdAttach,
	CmdLog,

	// service cmds
	CmdKill,
	CmdStatus,
	CmdStatistics,
}

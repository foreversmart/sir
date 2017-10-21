package commands

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CmdAttach = cli.Command{
	Name:      "attach",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "attach stdin & stdout to a task process",
	Action:    ActionAttach,
}

func ActionAttach(c *cli.Context) error {
	taskName := c.Args().First()
	// TODO check inputs

	// // 调用api
	// var response map[string]interface{}
	// httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/"+taskName+"/send"), message, &response)
	// // TODO handle error

	// println()
	// fmt.Println(utils.Style.Success("[INFO]"), "MESSAGE SENT TO TASK", utils.Style.Title(taskName), "SUCCESSFULLY", "\n")

	return nil
}

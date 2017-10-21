package commands

import (
	"fmt"
	"net/http"
	"sir/cli/config"
	"sir/cli/utils"
	"sir/lib/httpclient"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdSend = cli.Command{
	Name:      "send",
	UsageText: "<task> -m \"xxx\"",
	Category:  string(TaskCategory),
	Usage:     "send stdin to a task process",
	Action:    ActionSend,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "m",
			Usage: "must have, send message to stdin of task",
		},
	},
}

func ActionSend(c *cli.Context) error {
	taskName := c.Args().First()
	message := c.String("m")
	// TODO check inputs

	// 调用api
	var response map[string]interface{}
	httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/"+taskName+"/send"), message, &response)
	// TODO handle error

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "MESSAGE SENT TO TASK", utils.Style.Title(taskName), "SUCCESSFULLY", "\n")

	return nil
}

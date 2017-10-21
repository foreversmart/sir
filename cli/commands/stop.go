package commands

import (
	"fmt"
	"net/http"
	"sir/cli/config"
	"sir/cli/utils"
	"sir/lib/httpclient"

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
	taskName := c.Args().First()

	var response map[string]interface{}
	httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/"+taskName+"/stop"), nil, &response)

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "TASK", utils.Style.Title(taskName), "STOPED", "\n")

	return nil
}

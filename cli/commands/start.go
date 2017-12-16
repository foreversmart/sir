package commands

import (
	"fmt"
	"net/http"

	"github.com/foreversmart/sir/cli/config"
	"github.com/foreversmart/sir/cli/utils"
	"github.com/foreversmart/sir/lib/httpclient"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdStart = cli.Command{
	Name:      "start",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "start and daemonize a task",
	Action:    ActionStart,
}

func ActionStart(c *cli.Context) error {
	taskName := c.Args().First()

	var response map[string]interface{}
	httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/"+taskName+"/start"), nil, &response)

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "TASK", utils.Style.Title(taskName), "STARTED", "\n")

	return nil
}

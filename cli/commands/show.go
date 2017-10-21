package commands

import (
	"fmt"
	"net/http"
	"sir/cli/config"
	"sir/cli/utils"
	"sir/lib/httpclient"
	"sir/models"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdShow = cli.Command{
	Name:      "show",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "describe all parameters of a task",
	Action:    ActionShow,
}

func ActionShow(c *cli.Context) error {
	taskName := c.Args().First()
	// 调用api
	var response map[string]models.Task
	httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task/"+taskName), nil, &response)
	// TODO handle error

	fmt.Println(response)
	task := response["data"]

	utils.RenderTask(&task, c)

	return nil
}

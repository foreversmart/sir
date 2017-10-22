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

var CmdList = cli.Command{
	Name:      "list",
	UsageText: "",
	Category:  string(TaskCategory),
	Usage:     "list all tasks",
	Action:    ActionList,
}

func ActionList(c *cli.Context) error {

	var response map[string][]models.Task
	httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task"), nil, &response)

	// TODO handle error

	list := response["data"]

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "TASK LIST\n")

	utils.RenderTaskList(list, c)

	return nil
}

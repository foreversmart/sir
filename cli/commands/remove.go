package commands

import (
	"fmt"
	"net/http"

	"github.com/foreversmart/sir/cli/config"
	"github.com/foreversmart/sir/cli/utils"
	"github.com/foreversmart/sir/lib/httpclient"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdRemove = cli.Command{
	Name:      "remove",
	UsageText: "<task>",
	Category:  string(ConfigCategory),
	Usage:     "remove a task config by name",
	Action:    ActionRemove,
}

func ActionRemove(c *cli.Context) error {
	taskName := c.Args().First()
	// 调用api
	var response map[string]interface{}
	httpclient.Client.DoJSON(http.MethodDelete, config.ApiPath("/task/"+taskName), nil, &response)
	// TODO handle error

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "REMOVE TASK", utils.Style.Title(taskName), utils.Style.Success("SUCCESS"), "\n")

	return nil
}

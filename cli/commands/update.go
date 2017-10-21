package commands

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sir/cli/config"
	"sir/cli/utils"
	"sir/lib/httpclient"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdUpdate = cli.Command{
	Name:      "update",
	UsageText: "<task>",
	Category:  string(ConfigCategory),
	Usage:     "update task configs with default editor",
	Action:    ActionUpdate,
}

func ActionUpdate(c *cli.Context) error {
	taskName := c.Args().First()
	// 调用api
	var response map[string]string
	httpclient.Client.DoJSON(http.MethodPut, config.ApiPath("/task/"+taskName), nil, &response)
	// TODO handle error

	configPath := response["data"]

	editorPath, _ := exec.LookPath("vi")
	// TODO handle error

	cmd := exec.Command(editorPath, configPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "EDIT CONFIG DONE")
	fmt.Println(utils.Style.Success("[INFO]"), `PLEASE RESTART TASK PROCESS IF YOU WANT TO APPLY NEW CONFIG, BY:
       $ sir restart`, utils.Style.Title(taskName))

	return nil
}

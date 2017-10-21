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

var CmdLog = cli.Command{
	Name:      "log",
	UsageText: "[<task>]",
	Category:  string(TaskCategory),
	Usage:     "stream logs file. Default stream all tasks logs",
	Action:    ActionLog,
}

func ActionLog(c *cli.Context) error {
	taskName := c.Args().First()

	var response map[string]string

	if taskName == "" {
		httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task/"+taskName+"/log"), nil, &response)
	} else {
		httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task/log"), nil, &response)
	}

	// TODO handle error

	configPath := response["data"]

	editorPath, _ := exec.LookPath("tail")
	// TODO handle error

	cmd := exec.Command(editorPath, "-f", configPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "CHECK LOGS DONE")

	return nil
}

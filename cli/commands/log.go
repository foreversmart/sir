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
	UsageText: "<task> [--type \"err\"|\"std\"]",
	Category:  string(TaskCategory),
	Usage:     "stream task logs. default is std log",
	Action:    ActionLog,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "type",
			Usage: "log type, err | std, default is out",
		},
	},
}

func ActionLog(c *cli.Context) error {
	taskName := c.Args().First()
	logType := c.String("type")
	if logType == "" {
		logType = "std"
	}

	var response map[string]string

	httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task/"+taskName+"/log"), nil, &response)
	// TODO handle error

	path := response[logType]

	editorPath, _ := exec.LookPath("tail")
	// TODO handle error

	cmd := exec.Command(editorPath, "-f", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "CHECK LOGS DONE")

	return nil
}

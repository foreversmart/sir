package commands

import (
	"fmt"
	"net/http"

	"github.com/foreversmart/sir/cli/config"
	"github.com/foreversmart/sir/cli/utils"
	"github.com/foreversmart/sir/lib/httpclient"
	"github.com/foreversmart/sir/models"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdStatistics = cli.Command{
	Name:      "statistics",
	UsageText: "[<task>] [-from <fromtime>] [-to <endtime>]",
	Category:  string(ServiceCategory),
	Usage:     "show statistics, default sir daemon statistics",
	Action:    ActionStatistics,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Usage: "specify start time",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "specify end time",
		},
	},
}

func ActionStatistics(c *cli.Context) error {
	taskName := c.Args().First()
	from := c.String("from")
	to := c.String("to")

	query := "?from=" + from + "&to=" + to

	// 调用api
	var response map[string][]models.Statistics
	httpclient.Client.DoJSON(http.MethodGet, config.ApiPath("/task/"+taskName+"/statistics"+query), nil, &response)
	// TODO handle error

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "STATISTICS INFO", "\n")

	statistics := response["data"]
	utils.RenderStatistics(statistics, c)

	return nil
}

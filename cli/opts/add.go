package opts

import (
	"sir/models"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

type AddOpts struct {
	*models.TaskConfig
}

func ParseAddOpts(c *cli.Context) *AddOpts {

	// watch
	watch := false
	watchdir := c.String("watch")
	if watchdir != "" {
		watch = true
	}

	// env
	env := []string{}
	envstr := c.String("env")
	if envstr != "" {
		env = strings.Split(envstr, "&")
	}

	// TODO parse other params...

	return &AddOpts{
		TaskConfig: &models.TaskConfig{
			Cmd:        c.Args().First(),
			Name:       c.String("name"),
			Watch:      watch,
			WatchDir:   watchdir,
			Env:        env,
			Rules:      []*models.RuleConfig{},
			LogConfigs: &models.LogConfig{},
		},
	}
}

func (o *AddOpts) IsValid() bool {
	return o.Name != "" && o.Cmd != ""
}

package opts

import (
	"net/url"
	"sir/models"

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
	env := make(map[string]string)
	envstr := c.String("env")
	if envstr != "" {
		m, err := url.ParseQuery(envstr)
		if err != nil {
			return &AddOpts{}
		}
		for k, v := range (map[string][]string)(m) {
			env[k] = v[0]
		}
	}

	// TODO parse other params...

	return &AddOpts{
		TaskConfig: &models.TaskConfig{
			Cmd:      c.Args().First(),
			Name:     c.String("name"),
			Watch:    watch,
			WatchDir: watchdir,
			Env:      env,
		},
	}
}

func (o *AddOpts) IsValid() bool {
	return o.Name != "" && o.Cmd != ""
}

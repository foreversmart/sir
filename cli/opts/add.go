package opts

import cli "gopkg.in/urfave/cli.v1"

type AddOpts struct {
	Cmd  string
	Name string
}

func ParseAddOpts(c *cli.Context) *AddOpts {
	return &AddOpts{
		Cmd:  c.Args().First(),
		Name: c.String("name"),
	}
}

func (o *AddOpts) IsValid() bool {
	return o.Name != "" && o.Cmd != ""
}

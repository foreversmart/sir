package main

import (
	"os"

	"sir/cli/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	initApp().Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Sir"
	app.Usage = "May I help you?"
	app.Version = "0.0.1"
	app.UsageText = "sir [global options] command [command options] [arguments...]"
	app.CustomAppHelpTemplate = CustomAppHelpTemplate

	app.Commands = commands.CliCmds
	app.Action = func(c *cli.Context) error {
		// default action
		cli.ShowAppHelp(c)
		return nil
	}

	app.Before = func(c *cli.Context) error {
		// fmt.Println(chalk.Magenta.Color("Good to go~"))
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show help",
		},
	}
	app.HideHelp = true

	return app
}

const CustomAppHelpTemplate = `
   ___________ 
  / __/  _/ _ \ 
 _\ \_/ // , _/
/___/___/_/|_|    v{{.Version}}
							 

{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}


Usage: {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}


Commands: {{range .VisibleCategories}}{{if .Name}}

	> {{.Name}}:{{end}}{{range .VisibleCommands}}
		{{join .Names ", "}} {{.UsageText}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{if .VisibleFlags}}


Global Options:

	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}


Copyright:

{{.Copyright}}{{end}}
`

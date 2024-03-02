package main

import (
	"os"

	"github.com/enuesaa/loadii/pkg/command"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:    "loadii",
		Version: "0.0.1",
		Usage:   "A CLI tool to watch file changes and execute some operation.",
		Args:    true,
		Commands: []*cli.Command{
			command.NewExecCommand(repos),
			command.NewRunCommand(repos),
			command.NewServeCommand(repos),
		},
		Suggest: true,
	}

	// disable default
	app.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

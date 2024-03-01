package main

import (
	"log"
	"os"

	"github.com/enuesaa/loadii/pkg/command"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	app := &cli.App{
		Name:      "loadii",
		Version:   "0.0.2",
		Usage:     "Instant web server",
		Args:      true,
		Commands: []*cli.Command{
			&command.ServeCommand,
			&command.ExecCommand,
			&command.RunCommand,
		},
		Suggest:   true,
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
		log.Fatalf("Error: %s", err.Error())
	}
}

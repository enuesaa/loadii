package main

import (
	"log"
	"os"

	"github.com/enuesaa/loadii/internal/di"
	"github.com/enuesaa/loadii/internal/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	deps := di.New()

	var servePort int
	var serveDir string

	cmd := &cli.App{
		Name:  "loadii",
		Version: "0.0.11",
		Usage: "Instant web server for development purposes",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Port",
				Value:       3000,
				Destination: &servePort,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Usage:       "Serve dir",
				Value:       ".",
				Destination: &serveDir,
			},
		},
		Action: func(*cli.Context) error {
			usecases := usecase.New(deps)
			go usecases.Serve(serveDir, servePort)

			return usecases.Watch(serveDir)
		},
	}

	// disable default
	cmd.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
FLAGS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

	if err := cmd.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

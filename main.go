package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/loadii/internal/di"
	"github.com/enuesaa/loadii/internal/usecase"
	"github.com/urfave/cli/v2"
)

var servePort int
var serveDir string = "."

func main() {
	deps := di.New()

	cmd := &cli.App{
		Name:      "loadii",
		Version:   "0.0.12",
		Usage:     "Instant web server for development purposes",
		Args:      true,
		ArgsUsage: "<dir>",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Port",
				Value:       3000,
				Destination: &servePort,
			},
		},
		Action: func(c *cli.Context) error {
			args := c.Args()
			if args.Len() > 1 {
				return fmt.Errorf("Invalid arguments. Please specify exactly one argument.")
			}
			if args.Len() == 1 {
				serveDir = args.First()
			}
			usecases := usecase.New(deps)

			go func ()  {
				if err := usecases.Watch(serveDir); err != nil {
					deps.Log.Info("Error: %s", err.Error())
				}	
			}()
			return usecases.Serve(serveDir, servePort)
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

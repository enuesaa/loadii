package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

    "github.com/urfave/cli/v2"
	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/usecase"
)

func main() {
	repos := repository.New()

	var servePort int

	app := &cli.App{
        Name:  "loadii",
        Usage: "Instant web server for development",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Serve port",
				Value:       3000,
				Destination: &servePort,
			},
		},
        Action: func(*cli.Context) error {
			sigch := make(chan os.Signal, 1)
			signal.Notify(sigch, syscall.SIGTERM)

			go usecase.Serve(repos, sigch)

			return usecase.Watch(repos, ".")
        },
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
FLAGS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

    if err := app.Run(os.Args); err != nil {
		log.Panicf("Error: %s", err.Error())
    }
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	var servePort int
	var serveDir string

	app := &cli.App{
		Name:  "loadii",
		Version: "0.0.10",
		Usage: "Instant web server for development",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Serve port",
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
			sigch := make(chan os.Signal, 1)
			signal.Notify(sigch, syscall.SIGTERM)

			go usecase.Serve(repos, serveDir, sigch)

			return usecase.Watch(repos, serveDir)
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
		log.Fatalf("Error: %s", err.Error())
	}
}

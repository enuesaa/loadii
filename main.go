package main

import (
	"log"
	"os"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	var watchpath string
	var watchExcludePath string
	var workdir string
	var servepath string
	var serveport int

	app := &cli.App{
		Name:    "loadii",
		Version: "0.0.3",
		Usage:   "A CLI tool to watch file changes and execute a command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "include",
				Usage:       "add path to watch",
				Value:       "",
				Destination: &watchpath,
			},
			&cli.StringFlag{
				Name:        "exclude",
				Usage:       "remove path to watch",
				Value:       "",
				Destination: &watchExcludePath,
			},
			&cli.StringFlag{
				Name:        "serve",
				Usage:       "serve",
				Value:       "",
				Destination: &servepath,
			},
			&cli.IntFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "port",
				Value: 3000,
				Destination: &serveport,
			},
			&cli.StringFlag{
				Name: "workdir",
				Aliases: []string{"w"},
				Usage: "workdir",
				Value: ".",
				Destination: &workdir,
			},
		},
		Args:      true,
		ArgsUsage: "commands",
		Action: func(c *cli.Context) error {
			if len(c.FlagNames()) == 0 {
				return cli.ShowAppHelp(c)
			}

			usecase.Watch(repos, watchpath)

			return nil
		},
		Suggest: true,
	}

	// disable default
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		return err
	}
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

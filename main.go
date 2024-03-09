package main

import (
	"log"
	"os"

	"github.com/enuesaa/loadii/pkg/command"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	var watchpath string

	app := &cli.App{
		Name:    "loadii",
		Version: "0.0.2",
		Usage:   "A CLI tool to watch file changes and execute a task",
		Commands: []*cli.Command{
			command.NewExecCommand(repos, watchpath),
			command.NewRunCommand(repos, watchpath),
			command.NewServeCommand(repos, watchpath),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "watch",
				Aliases: []string{"w"},
				Usage:   "watch dir",
				Destination: &watchpath,
			},
		},
		Action: func(c *cli.Context) error {
			if watchpath == "" {
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

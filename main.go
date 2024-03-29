package main

import (
	"log"
	"os"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/enuesaa/loadii/pkg/watch"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	var watchIncludes cli.StringSlice
	var watchExcludes cli.StringSlice
	var workdir string
	var servePath string
	var servePort int
	var autoApprove bool

	app := &cli.App{
		Name:    "loadii",
		Version: "0.0.4",
		Usage:   "A CLI tool to watch file changes and execute a command",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "yes",
				Aliases:     []string{"y"},
				Value:       false,
				Usage:       "Approve command execution",
				Destination: &autoApprove,
			},
			&cli.StringSliceFlag{
				Name:        "include",
				Usage:       "Add path to watch",
				Value:       cli.NewStringSlice("."),
				Destination: &watchIncludes,
				Category:    "watch",
			},
			&cli.StringSliceFlag{
				Name:        "exclude",
				Usage:       "Remove path to watch",
				Destination: &watchExcludes,
				Category:    "watch",
			},
			&cli.StringFlag{
				Name:        "serve",
				Usage:       "Serve dir",
				Destination: &servePath,
				Category:    "serve",
			},
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Serve port",
				Value:       3000,
				Destination: &servePort,
				Category:    "serve",
			},
			&cli.StringFlag{
				Name:        "workdir",
				Aliases:     []string{"w"},
				Usage:       "Command execution dir",
				Value:       ".",
				Destination: &workdir,
				Category:    "serve",
			},
		},
		Args:      true,
		ArgsUsage: "commands",
		Action: func(c *cli.Context) error {
			commands := c.Args().Slice()

			// When no flag, no args passed
			if len(c.FlagNames()) == 0 && len(commands) == 0 {
				return cli.ShowAppHelp(c)
			}

			usecase.Plan(repos, usecase.PlanProps{
				ServePath: servePath,
				ServePort: servePort,
				Commands: commands,
				WatchIncludes: watchIncludes.Value(),
				WatchExcludes: watchExcludes.Value(),
			})

			if len(commands) > 0 {
				if err := usecase.Confirm(repos, commands, autoApprove); err != nil {
					return err
				}
				usecase.Exec(repos, workdir, commands)
			}

			if servePath != "" {
				go usecase.Serve(repos, servePath, servePort)
			}

			options := []watch.Option{
				watch.WithIncludes(watchIncludes.Value()),
				watch.WithExcludes(watchExcludes.Value()),
			}
			if len(commands) > 0 {
				options = append(options, watch.WithCallback(func() {
					usecase.Exec(repos, workdir, commands)
				}))
			}

			return usecase.Watch(repos, options...)
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
	{{.HelpName}} {{if .VisibleFlags}}[flags]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
FLAGS:{{range .VisibleFlagCategories}}{{if len .Name}}  [{{.Name}}]{{end}}
	{{range .Flags}}{{.}}
	{{end}}
{{end}}{{end}}`

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

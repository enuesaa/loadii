package main

import (
	"os"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:    "loadii",
		Version: "0.0.7",
		Usage:   "A CLI tool to watch file changes and execute a command",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "yes",
				Aliases: []string{"y"},
				Value:   false,
				Usage:   "Approve command execution",
			},
			&cli.StringSliceFlag{
				Name:     "include",
				Usage:    "Add path to watch",
				Value:    cli.NewStringSlice("."),
				Category: "WATCH",
			},
			&cli.StringSliceFlag{
				Name:     "exclude",
				Usage:    "Remove path to watch",
				Category: "WATCH",
			},
			&cli.StringFlag{
				Name:     "serve",
				Usage:    "Serve dir",
				Category: "SERVE",
			},
			&cli.IntFlag{
				Name:     "port",
				Usage:    "Serve port",
				Value:    3000,
				Category: "SERVE",
			},
			&cli.StringFlag{
				Name:     "workdir",
				Aliases:  []string{"w"},
				Usage:    "Command execution dir",
				Value:    ".",
				Category: "SERVE",
			},
			&cli.StringFlag{
				Name:     "go",
				Usage:    "",
				Required: false,
				Value: ".",
				DefaultText: ".",
				Category: "SERVE",
			},
		},
		Args:      true,
		ArgsUsage: "commands",
		Action: func(c *cli.Context) error {
			plan := usecase.Plan{
				Workdir:       c.String("workdir"),
				ServePath:     c.String("serve"),
				ServePort:     c.Int("port"),
				Commands:      c.Args().Slice(),
				WatchIncludes: c.StringSlice("include"),
				WatchExcludes: c.StringSlice("exclude"),
			}
			autoApprove := c.Bool("yes")

			// When no flag, no args passed
			if len(c.FlagNames()) == 0 && len(plan.Commands) == 0 {
				return cli.ShowAppHelp(c)
			}

			plan.Print(repos)
			if !autoApprove {
				if err := plan.Confirm(repos); err != nil {
					return err
				}
			}

			if len(plan.Commands) > 0 {
				usecase.Exec(repos, plan)
			}
			if plan.ServePath != "" {
				go usecase.Serve(repos, plan)
			}

			return usecase.Watch(repos, plan)
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
FLAGS:{{range .VisibleFlagCategories}}{{if len .Name}}  {{.Name}}{{end}}
	{{range .Flags}}{{.}}
	{{end}}
{{end}}{{end}}`

	if err := app.Run(os.Args); err != nil {
		repos.Log.Fatal(err)
	}
}

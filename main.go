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
		Version: "0.0.3",
		Usage:   "A CLI tool to watch file changes and execute a task",
		Commands: []*cli.Command{
			// loadii go run .
			// loadii --watch . go run .
			command.NewExecCommand(repos, &watchpath),
			//Deprecated
			// 他では見かけないコマンドなので、exec に比べて認知負荷が高い
			// 例えば go run . をこれで実現するには loadii run main.go を渡す必要があるが、全く別の見た目で大きな隔たりがある
			command.NewRunCommand(repos, &watchpath),
			// loadii --serve ./dist 
			// loadii --watch . --serve ./dist pnpm build
			// loadii exec go run . | loadii serve ./dist
			command.NewServeCommand(repos, &watchpath),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "watch",
				Aliases:     []string{"w"},
				Usage:       "watch dir",
				Value:       ".",
				Destination: &watchpath,
			},
		},
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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/tryserve/pkg/repository"
	"github.com/enuesaa/tryserve/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	watchmode := false

	app := &cli.App{
		Name:      "tryserve",
		Version:   "0.0.2",
		Usage:     "Instant web server",
		Args:      true,
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "watch",
				Value:       false,
				Usage:       "run watch mode",
				Destination: &watchmode,
			},
		},
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				return cli.ShowAppHelp(c)
			}
			if watchmode {
				fmt.Printf("running on watch mode\n")
			}

			repos := repository.New()
			ext := repos.Fs.Ext(path)
			// TODO: change this logic. if path is file, run app.
			if ext == "" {
				return usecase.Serve(repos, path)
			}
			return usecase.RunApp(repos, path)
		},
	}

	// disable default
	app.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}} {{if .VisibleFlags}}[global options]{{end}}
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
		log.Fatal(err)
	}
}

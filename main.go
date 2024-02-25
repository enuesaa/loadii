package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/tryup/pkg/repository"
	"github.com/enuesaa/tryup/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	watchmode := false
	port := 3000

	app := &cli.App{
		Name:      "tryup",
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
			&cli.IntFlag{
				Name:        "port",
				Value:       3000,
				Usage:       "port",
				Destination: &port,
				Action: func(ctx *cli.Context, v int) error {
					// see https://cli.urfave.org/v2/examples/flags/
					if v > 65537 {
						return fmt.Errorf("invalid value %d passed to flag --port", v)
					}
					return nil
				},
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
			if !repos.Fs.IsExist(path) {
				return fmt.Errorf("not found: %s", path)
			}

			isDir, err := repos.Fs.IsDir(path)
			if err != nil {
				return fmt.Errorf("unknown error occuerd: %s", err.Error())
			}
			if isDir {
				return usecase.Serve(repos, path, port)
			}
			return usecase.RunApp(repos, path)
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
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package command

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

var ServeCommand = cli.Command{
	Name:    "serve",
	Usage:   "serve instant web server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "workdir",
			Aliases: []string{"w"},
			Value:       "./",
			Usage:       "workdir",
		},
		&cli.IntFlag{
			Name:        "port",
			Value:       3000,
			Usage:       "port",
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
		workdir := c.String("workdir")
		port := c.Int("port")

		repos := repository.New()
		usecase.Watch(repos)

		return usecase.Serve(repos, workdir, port)
	},
}

package command

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewServeCommand(repos repository.Repos, watchpath *string) *cli.Command {
	var port int

	cmd := cli.Command{
		Name:  "serve",
		Usage: "serve instant web server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Value:       3000,
				Usage:       "port",
				Destination: &port,
			},
		},
		Args:      true,
		ArgsUsage: "<path>",
		Action: func(c *cli.Context) error {
			path := "."
			if c.Args().Len() > 0 {
				path = c.Args().Get(0)
			}

			return usecase.ServeWatch(repos, *watchpath, path, port)
		},
	}

	return &cmd
}

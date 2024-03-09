package command

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewServeCommand(repos repository.Repos) *cli.Command {
	cmd := cli.Command{
		Name:  "serve",
		Usage: "serve instant web server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Value: 3000,
				Usage: "port",
			},
		},
		Args: true,
		ArgsUsage: "path",
		Action: func(c *cli.Context) error {
			path := "."
			if c.Args().Len() > 0 {
				path = c.Args().Get(0)
			}
			port := c.Int("port")

			return usecase.ServeWatch(repos, path, port)
		},
	}

	return &cmd
}

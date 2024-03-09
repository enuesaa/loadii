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
			&cli.StringFlag{
				Name:    "workdir",
				Value:   "./",
				Usage:   "workdir",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 3000,
				Usage: "port",
			},
		},
		Action: func(c *cli.Context) error {
			workdir := c.String("workdir")
			port := c.Int("port")

			return usecase.ServeWatch(repos, workdir, port)
		},
	}

	return &cmd
}

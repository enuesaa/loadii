package command

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

var RunCommand = cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "run app. Currently, this command supports golang app.",
	Action: func(c *cli.Context) error {
		path := c.Args().Get(0)
		if path == "" {
			return cli.ShowAppHelp(c)
		}

		repos := repository.New()

		return usecase.RunApp(repos, path)
	},
}

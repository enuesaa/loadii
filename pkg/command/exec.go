package command

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

var ExecCommand = cli.Command{
	Name:    "exec",
	Aliases: []string{"e"},
	Usage:   "exec commands",
	Args: true,
	ArgsUsage: "commands",
	Action: func(c *cli.Context) error {
		commands := c.Args().Slice()
		repos := repository.New()

		if err := usecase.Exec(commands); err != nil {
			return err
		}
		callback := func () {
			usecase.Exec(commands)
		}
		if err := usecase.Watch(repos, &callback); err != nil {
			return err
		}
		usecase.Sleep()

		return nil
	},
}

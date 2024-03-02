package command

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

var ExecCommand = cli.Command{
	Name:    "exec",
	Usage:   "exec commands",
	Args: true,
	ArgsUsage: "commands",
	Before: func(c *cli.Context) error {
		commands := c.Args().Slice()
		if len(commands) == 0 {
			return fmt.Errorf("please specify command")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		commands := c.Args().Slice()
		repos := repository.New()

		return usecase.ExecWatch(repos, commands)
	},
}

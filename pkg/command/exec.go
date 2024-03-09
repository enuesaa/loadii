package command

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewExecCommand(repos repository.Repos, watchpath *string) *cli.Command {
	cmd := cli.Command{
		Name:      "exec",
		Usage:     "exec commands",
		Args:      true,
		ArgsUsage: "commands",
		Before: func(c *cli.Context) error {
			commands := c.Args().Slice()
			if len(commands) == 0 {
				return fmt.Errorf("Required arguments are missing.\nPlease pass a command like `loadii exec echo a`.\n")
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			commands := c.Args().Slice()

			return usecase.ExecWatch(repos, *watchpath, commands)
		},
	}

	return &cmd
}

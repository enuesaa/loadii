package command

import (
	"fmt"

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
		fmt.Printf("exec command: %v\n", commands)

		repos := repository.New()

		if err := usecase.Exec(commands); err != nil {
			return err
		}

		if err := usecase.Watch(repos); err != nil {
			return err
		}
		return nil
	},
}

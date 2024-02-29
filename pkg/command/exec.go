package command

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

var ExecCommand = cli.Command{
	Name:    "exec",
	Aliases: []string{"e"},
	Usage:   "exec commands",
	Action: func(c *cli.Context) error {
		commands := c.Args().Slice()

		fmt.Printf("exec command: %v\n", commands)

		return usecase.Exec(commands[0], commands[1:])
	},
}

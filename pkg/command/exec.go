package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ExecCommand = cli.Command{
	Name:    "exec",
	Aliases: []string{"e"},
	Usage:   "exec commands",
	Action: func(c *cli.Context) error {
		fmt.Println("exec command")
		return nil
	},
}

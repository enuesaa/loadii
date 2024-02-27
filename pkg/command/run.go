package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var RunCommand = cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "run app. Currently, this command supports golang app.",
	Action: func(c *cli.Context) error {
		fmt.Println("run")
		return nil
	},
}

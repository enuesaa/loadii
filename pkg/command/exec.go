package command

import (
	"fmt"
	"time"

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

		if err := usecase.Exec(commands); err != nil {
			return err
		}

		if err := usecase.Watch("./"); err != nil {
			return err
		}
		for range 10 {
			time.Sleep(1 * time.Second)
			fmt.Printf("a\n")
		}
		return nil
	},
}

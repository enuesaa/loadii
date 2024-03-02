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
	Action: func(c *cli.Context) error {
		commands := c.Args().Slice()
		if len(commands) == 0 {
			return fmt.Errorf("please specify command")
		}
		repos := repository.New()

		if err := usecase.ExecWatch(repos, commands); err != nil {
			return err
		}
		<-make(chan struct{})

		return nil
	},
}

package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, plan Plan) error {
	if len(plan.Commands) == 0 {
		return fmt.Errorf("Failed to run the command ``")
	}

	execctl := exec.New(repos)
	execctl.Workdir = plan.Workdir
	execctl.Command = plan.Commands[0]
	if len(plan.Commands) > 1 {
		execctl.Args = plan.Commands[1:]
	}

	return execctl.Exec()
}

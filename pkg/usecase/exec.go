package usecase

import (
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, plan Plan) error {
	execctl := exec.New(repos)

	if err := execctl.Exec(plan.Workdir, plan.Commands); err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, workdir string, commands []string) error {
	execctl := exec.New(repos)

	if err := execctl.Exec(workdir, commands); err != nil {
		return err
	}
	return nil
}

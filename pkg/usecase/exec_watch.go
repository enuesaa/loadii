package usecase

import (
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func ExecWatch(repos repository.Repos, includes []string, excludes []string, commands []string, workdir string) error {
	execctl := exec.New(repos)

	if err := execctl.Exec(workdir, commands); err != nil {
		return err
	}

	watchctl := watch.New(repos)
	watchctl.Includes = includes
	watchctl.Excludes = excludes
	defer watchctl.Close()

	watchctl.AddCallback(func() {
		execctl.Exec(workdir, commands)
	})

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}
	watchctl.Wait()

	return nil
}

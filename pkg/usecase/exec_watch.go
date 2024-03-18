package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

// TODO change argument. use struct
func ExecWatch(repos repository.Repos, includes []string, excludes []string, commands []string, workdir string) error {
	if err := Exec(repos, commands, workdir); err != nil {
		return err
	}

	watchctl := watch.New(repos)
	watchctl.Includes = includes
	watchctl.Excludes = excludes
	defer watchctl.Close()

	watchctl.AddCallback(func() {
		Exec(repos, commands, workdir)
	})

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}
	watchctl.Wait()

	return nil
}

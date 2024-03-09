package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func ExecWatch(repos repository.Repos, watchpath string, commands []string) error {
	if err := Exec(repos, commands); err != nil {
		return err
	}

	watchctl := watch.New(repos)
	watchctl.WatchPath = watchpath
	defer watchctl.Close()

	watchctl.AddCallback(func() {
		Exec(repos, commands)
	})

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}
	watchctl.Wait()

	return nil
}

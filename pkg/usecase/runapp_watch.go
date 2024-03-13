package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func RunAppWatch(repos repository.Repos, watchpath string, path string, args []string) error {
	watchctl := watch.New(repos)
	watchctl.WatchPath = watchpath
	defer watchctl.Close()

	watchctl.AddCallback(func() {
		RunApp(repos, path, args)
	})

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}
	if err := RunApp(repos, path, args); err != nil {
		return err
	}
	watchctl.Wait()

	return nil
}

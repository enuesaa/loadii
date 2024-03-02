package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func ExecWatch(repos repository.Repos, commands []string) error {
	if err := Exec(repos, commands); err != nil {
		return err
	}

	watchctl := watch.New(repos)
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

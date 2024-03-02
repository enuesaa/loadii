package usecase

import (
	"log"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func RunAppWatch(repos repository.Repos, path string) error {
	watchctl := watch.New(repos)
	defer watchctl.Close()

	watchctl.AddCallback(func() {
		RunApp(repos, path)
	})

	if err := watchctl.Watch(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := RunApp(repos, path); err != nil {
		return err
	}
	watchctl.Wait()

	return nil
}

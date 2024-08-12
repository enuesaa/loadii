package usecase

import (
	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/watch"
)

func Watch(repos repository.Repos, include string) error {
	watchctl := watch.New(repos)
	watchctl.Includes = []string{include}
	watchctl.Excludes = []string{}
	defer watchctl.Close()

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}

	<-make(chan struct{})

	return nil
}

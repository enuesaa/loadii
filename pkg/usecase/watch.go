package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos, includes []string, excludes []string) error {
	watchctl := watch.New(repos)
	defer watchctl.Close()

	watchctl.Includes = includes
	watchctl.Excludes = excludes

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}

	<-make(chan struct{})

	return nil
}

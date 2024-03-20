package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos, options ...watch.Option) error {
	watchctl := watch.New(repos, options...)
	defer watchctl.Close()

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}

	<-make(chan struct{})

	return nil
}

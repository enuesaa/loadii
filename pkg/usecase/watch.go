package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos) {
	go func() {
		watchctl := watch.New(repos)
		defer watchctl.Close()

		if err := watchctl.Watch(); err != nil {
			repos.Log.Fatal(err)
		}

		<-make(chan struct{})
	}()
}

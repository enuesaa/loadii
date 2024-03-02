package usecase

import (
	"log"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos) {
	go func() {
		watchctl := watch.New(repos)
		defer watchctl.Close()

		if err := watchctl.Watch(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		<-make(chan struct{})
	}()
}

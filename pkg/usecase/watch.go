package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos) error {
	watchctl := watch.New(repos)
	defer watchctl.Close()

	fmt.Printf("watching ./\n")

	return watchctl.Watch()
}

func WatchSleep(repos repository.Repos) error {
	watchctl := watch.New(repos)
	defer watchctl.Close()

	fmt.Printf("watching ./\n")

	if err := watchctl.Watch(); err != nil {
		return err
	}

	<-make(chan struct{})

	return nil
}

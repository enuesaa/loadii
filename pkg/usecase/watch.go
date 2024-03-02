package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos, callback *func ()) error {
	watchctl := watch.New(repos)
	// defer watchctl.Close()

	fmt.Printf("watching ./\n")

	if callback != nil {
		watchctl.Callback = callback
	}

	return watchctl.Watch()
}

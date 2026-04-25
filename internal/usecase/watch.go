package usecase

import (
	"github.com/enuesaa/loadii/internal/deps"
	"github.com/enuesaa/loadii/internal/watch"
)

func Watch(repos *deps.Repos, include string) error {
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

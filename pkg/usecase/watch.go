package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos) error {
	watchctl := watch.New(repos)
	return watchctl.Watch()

}

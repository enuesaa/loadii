package watch

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/fsnotify/fsnotify"
)

func New(repos repository.Repos) Watchctl {
	return Watchctl{
		repos:     repos,
		callbacks: make([]func(), 0),
	}
}

type Watchctl struct {
	repos     repository.Repos
	watcher   *fsnotify.Watcher
	callbacks []func()
}

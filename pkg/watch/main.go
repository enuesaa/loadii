package watch

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/fsnotify/fsnotify"
)

func New(repos repository.Repos) Watchctl {
	return Watchctl {
		repos: repos,
		Callback: nil,
	}
}

type Watchctl struct {
	repos repository.Repos
	watcher *fsnotify.Watcher
	Callback *func()
}

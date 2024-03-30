package watch

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/fsnotify/fsnotify"
)

func New(repos repository.Repos, options ...Option) Watchctl {
	ctl := Watchctl{
		repos:   repos,
		options: Options{},
	}
	for _, fn := range options {
		fn(&ctl.options)
	}

	return ctl
}

type Watchctl struct {
	repos   repository.Repos
	watcher *fsnotify.Watcher
	options Options
}

// finally, call this function
func (ctl *Watchctl) Close() error {
	if ctl.watcher == nil {
		return nil
	}

	return ctl.watcher.Close()
}

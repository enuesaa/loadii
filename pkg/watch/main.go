package watch

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/fsnotify/fsnotify"
)

func New(repos repository.Repos) Watchctl {
	ctl := Watchctl{
		repos:     repos,
		Includes:  []string{},
		Excludes:  []string{},
		callbacks: []func(){},
	}
	return ctl
}

type Watchctl struct {
	repos     repository.Repos
	watcher   *fsnotify.Watcher
	Includes  []string
	Excludes  []string
	callbacks []func()
}

// finally, call this function
func (ctl *Watchctl) Close() error {
	if ctl.watcher == nil {
		return nil
	}

	return ctl.watcher.Close()
}

func (ctl *Watchctl) SetCallback(fn func()) {
	ctl.callbacks = []func(){fn}
}

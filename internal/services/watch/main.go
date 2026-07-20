package watch

import (
	"github.com/enuesaa/loadii/internal/di"
	"github.com/fsnotify/fsnotify"
)

func New(deps *di.Deps) Watchctl {
	ctl := Watchctl{
		running:   false,
		deps:      deps,
		Includes:  []string{},
		Excludes:  []string{},
		callbacks: []func(){},
	}
	return ctl
}

type Watchctl struct {
	running   bool
	deps      *di.Deps
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

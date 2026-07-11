package watch

import (
	"github.com/enuesaa/loadii/internal/deps"
	"github.com/fsnotify/fsnotify"
)

func New(repos *deps.Repos) Watchctl {
	ctl := Watchctl{
		running:   false,
		repos:     repos,
		Includes:  []string{},
		Excludes:  []string{},
		callbacks: []func(){},
	}
	return ctl
}

type Watchctl struct {
	running   bool
	repos     *deps.Repos
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

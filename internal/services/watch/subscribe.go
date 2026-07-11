package watch

import (
	"time"

	"github.com/fsnotify/fsnotify"
)

func (ctl *Watchctl) subscribe() {
	for {
		event, ok := <-ctl.watcher.Events
		if !ok {
			return
		}
		if event.Has(fsnotify.Write) {
			ctl.repos.Log.Info("file.modified %s", event.Name)
			ctl.triggerCallbacks()
		} else if event.Has(fsnotify.Remove) {
			ctl.repos.Log.Info("file.removed %s", event.Name)
			ctl.triggerCallbacks()
		} else if event.Has(fsnotify.Create) {
			ctl.repos.Log.Info("file.created %s", event.Name)
			ctl.triggerCallbacks()
		} else if event.Has(fsnotify.Rename) {
			// this event is equal to file.removed
			ctl.repos.Log.Info("file.removed %s", event.Name)
			ctl.triggerCallbacks()
		}
	}
}

func (ctl *Watchctl) triggerCallbacks() {
	// do not run on same time. As far as possible, this should be `Exactly Once With No Guarantee`.
	if ctl.running {
		return
	}
	ctl.running = true
	for _, fnc := range ctl.callbacks {
		fnc()
	}
	go func() {
		time.Sleep(10 * time.Second)
		ctl.running = false
	}()
}

package watch

import (
	"github.com/fsnotify/fsnotify"
)

func (ctl *Watchctl) Watch() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	ctl.watcher = watcher

	go func() {
		for {
			event, ok := <-ctl.watcher.Events
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				ctl.repos.Log.Info("modified: %s", event.Name)
				ctl.triggerCallbacks()
			} else if event.Has(fsnotify.Remove) {
				ctl.repos.Log.Info("deleted: %s", event.Name)
				ctl.triggerCallbacks()
			} else if event.Has(fsnotify.Create) {
				ctl.repos.Log.Info("created: %s", event.Name)
				ctl.triggerCallbacks()
			} else if event.Has(fsnotify.Rename) {
				// this seems deleted file.
				ctl.repos.Log.Info("deleted: %s", event.Name)
				ctl.triggerCallbacks()
			}
		}
	}()

	return ctl.watcher.Add("./")
}

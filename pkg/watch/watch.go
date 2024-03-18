package watch

import (
	"os"
	"path/filepath"
	"strings"

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

	for _, path := range ctl.Includes {
		if err := ctl.watcher.Add(path); err != nil {
			return err
		}
		err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				return nil
			}
			if strings.HasPrefix(path, ".") {
				return nil
			}

			return ctl.watcher.Add(path)
		})
		if err != nil {
			return err
		}
	}

	for _, path := range ctl.Excludes {
		if err := ctl.watcher.Remove(path); err != nil {
			return err
		}
		err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				return nil
			}
			if strings.HasPrefix(path, ".") {
				return nil
			}

			return ctl.watcher.Remove(path)
		})
		if err != nil {
			return err
		}
	}

	return nil
}

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

	go ctl.subscribe()

	ctl.setupIncludes()
	ctl.setupExcludes()

	return nil
}

func (ctl *Watchctl) setupIncludes() error {
	for _, path := range ctl.Includes {
		if err := ctl.watcher.Add(path); err != nil {
			return err
		}
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				return nil
			}
			// do not include dotfiles
			// this behavior should be configured.
			if strings.HasPrefix(path, ".") {
				return nil
			}

			return ctl.watcher.Add(path)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctl *Watchctl) setupExcludes() error {
	for _, path := range ctl.Excludes {
		if err := ctl.watcher.Remove(path); err != nil {
			return err
		}
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
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

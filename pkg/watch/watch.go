package watch

import (
	"log"

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
			event, ok := <-watcher.Events
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				log.Println("modified:", event.Name)
			} else if event.Has(fsnotify.Remove) {
				log.Println("deleted:", event.Name)
			} else if event.Has(fsnotify.Create) {
				log.Println("created:", event.Name)
			} else if event.Has(fsnotify.Rename) {
				// this seems deleted file.
				log.Println("deleted:", event.Name)
			}
		}
	}()

	go func() {
		for {
			err, ok := <-watcher.Errors
			if !ok {
				return
			}
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	return watcher.Add("./")
}

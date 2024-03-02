package watch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

// observer pattern みたいな機構がいい
func (ctl *Watchctl) trigger() {
	fmt.Printf("triggered!\n")
}

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
				log.Println("modified:", event.Name)
				ctl.trigger()
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
			err, ok := <-ctl.watcher.Errors
			if !ok {
				return
			}
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	return ctl.watcher.Add("./")
}

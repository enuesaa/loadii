package usecase

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	// TODO
	// defer watcher.Close()

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
	if err := watcher.Add("./"); err != nil {
		return err
	}
	for _, file := range watcher.WatchList() {
		fmt.Printf("watching: %s\n", file)
	}

	return nil
}

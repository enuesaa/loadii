package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer watcher.Close()

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

	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}

	for range 10 {
		time.Sleep(1 * time.Second)
		fmt.Printf("a\n")
	}
}

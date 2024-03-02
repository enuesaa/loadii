package usecase

import (
	"log"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func ExecWatch(repos repository.Repos, commands []string) error {
	if err := Exec(repos, commands); err != nil {
		return err
	}

	go func() {
		watchctl := watch.New(repos)
		defer watchctl.Close()

		if err := watchctl.Watch(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}

		callback := func ()  {
			Exec(repos, commands)
		}
		watchctl.Callback = &callback

		<-make(chan struct{})
	}()

	return nil
}

package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

//TODO: watch path should be configured in global flag `-w`
func ServeWatch(repos repository.Repos, basepath string, port int) error {
	go func() {
		watchctl := watch.New(repos)
		defer watchctl.Close()

		watchctl.WatchPath = basepath

		if err := watchctl.Watch(); err != nil {
			repos.Log.Fatal(err)
		}

		<-make(chan struct{})
	}()

	return Serve(repos, basepath, port)
}

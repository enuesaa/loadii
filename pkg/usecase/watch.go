package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/watch"
)

func Watch(repos repository.Repos, plan Plan) error {
	watchctl := watch.New(repos)
	watchctl.Includes = plan.WatchIncludes
	watchctl.Excludes = plan.WatchExcludes

	if len(plan.Commands) > 0 {
		watchctl.SetCallback(func() {
			Exec(repos, plan)
		})
	}
	defer watchctl.Close()

	if err := watchctl.Watch(); err != nil {
		repos.Log.Fatal(err)
	}

	<-make(chan struct{})

	return nil
}

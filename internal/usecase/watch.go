package usecase

import "github.com/enuesaa/loadii/internal/services/watch"

func (u *useCaseImpl) Watch(include string) error {
	watchctl := watch.New(u.deps)
	watchctl.Includes = []string{include}
	watchctl.Excludes = []string{}
	defer watchctl.Close()

	if err := watchctl.Watch(); err != nil {
		u.deps.Log.Fatal(err)
	}

	<-make(chan struct{})

	return nil
}

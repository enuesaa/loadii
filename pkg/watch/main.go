package watch

import "github.com/enuesaa/loadii/pkg/repository"

func New(repos repository.Repos) Watchctl {
	return Watchctl {
		repos: repos,
	}
}

type Watchctl struct {
	repos repository.Repos
}

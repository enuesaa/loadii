package exec

import (
	"github.com/enuesaa/loadii/pkg/repository"
)

func New(repos repository.Repos) Execctl {
	return Execctl{
		repos: repos,
	}
}

type Execctl struct {
	repos repository.Repos
}

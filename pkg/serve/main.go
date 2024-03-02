package serve

import "github.com/enuesaa/loadii/pkg/repository"

func New(repos repository.Repos) Servectl {
	return Servectl{
		Port:     3000,
		Basepath: ".",
	}
}

type Servectl struct {
	repos    repository.Repos
	Port     int
	Basepath string
}

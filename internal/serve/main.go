package serve

import "github.com/enuesaa/loadii/internal/repository"

func New(repos repository.Repos) Servectl {
	return Servectl{
		repos:    repos,
		Port:     3000,
		Basepath: ".",
	}
}

type Servectl struct {
	repos    repository.Repos
	Port     int
	Basepath string
}

package serve

import "github.com/enuesaa/loadii/internal/deps"

func New(repos *deps.Container) Servectl {
	return Servectl{
		repos:    repos,
		Port:     3000,
		Basepath: ".",
	}
}

type Servectl struct {
	repos    *deps.Container
	Port     int
	Basepath string
}

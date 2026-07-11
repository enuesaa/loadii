package serve

import "github.com/enuesaa/loadii/internal/di"

func New(deps *di.Deps) Servectl {
	return Servectl{
		repos:    deps,
		Port:     3000,
		Basepath: ".",
	}
}

type Servectl struct {
	repos    *di.Deps
	Port     int
	Basepath string
}

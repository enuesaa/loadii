package usecase

import "github.com/enuesaa/loadii/internal/di"

type UseCases interface {
	Watch(include string) error
	Serve(dir string, port int)
}

func New(deps *di.Deps) UseCases {
	return &useCaseImpl{
		deps: deps,
	}
}

type useCaseImpl struct {
	deps *di.Deps
}

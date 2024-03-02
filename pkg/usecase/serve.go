package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
)

func Serve(repos repository.Repos, basepath string, port int) error {
	servectl := serve.New(repos)
	servectl.Port = port
	servectl.Basepath = basepath

	return servectl.Listen()
}

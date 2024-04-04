package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
)

func Serve(repos repository.Repos, plan Plan) error {
	servectl := serve.New(repos)
	servectl.Port = plan.ServePort
	servectl.Basepath = plan.ServePath

	repos.Log.Info("Listening on %s", servectl.Addr())

	return servectl.Listen()
}

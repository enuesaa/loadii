package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
)

func Serve(repos repository.Repos, basepath string, port int) error {
	servectl := serve.New(repos)
	servectl.Port = port
	servectl.Basepath = basepath

	fmt.Printf("Listening on %s\n", servectl.Addr())

	return servectl.Listen()
}

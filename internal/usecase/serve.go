package usecase

import (
	"fmt"
	"os"

	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/serve"
)

func Serve(repos repository.Repos, dir string, port int, sigch chan os.Signal) {
	servectl := serve.New(repos)
	servectl.Port = port
	servectl.Basepath = dir

	fmt.Printf("Listening on %s\n", servectl.Addr())

	if err := servectl.Listen(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

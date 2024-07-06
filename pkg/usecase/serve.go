package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/cli"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
)

func Serve(repos repository.Repos) {
	servectl := serve.New(repos)
	servectl.Port = 3000
	servectl.Basepath = cli.ServeFlag.Workdir()

	fmt.Printf("Listening on %s", servectl.Addr())

	if err := servectl.Listen(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	// sig := <-sigCh
	// fmt.Printf("Received: %v\n", sig)
	// if err := execctl.Kill(); err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// }
}

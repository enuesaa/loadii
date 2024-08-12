package usecase

import (
	"fmt"
	"os"

	"github.com/enuesaa/loadii/internal/cli"
	"github.com/enuesaa/loadii/internal/exec"
	"github.com/enuesaa/loadii/internal/repository"
)

func ExecGo(repos repository.Repos, sigch chan os.Signal) {
	execctl := exec.New(repos)
	// TODO: 直感的でない
	execctl.Workdir = cli.GoFlag.Workdir()
	execctl.Command = "go"
	execctl.Args = []string{"run"}
	execctl.Args = append(execctl.Args, cli.GoFlag.Values()...)

	if err := execctl.Exec(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	sig := <-sigch
	fmt.Printf("Received: %v\n", sig)
	if err := execctl.Kill(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

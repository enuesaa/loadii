package usecase

import (
	"fmt"
	"os"

	"github.com/enuesaa/loadii/internal/cli"
	"github.com/enuesaa/loadii/internal/exec"
	"github.com/enuesaa/loadii/internal/repository"
)

func ExecPnpm(repos repository.Repos, sigch chan os.Signal) {
	execctl := exec.New(repos)
	execctl.Workdir = cli.PnpmFlag.Workdir()
	execctl.Command = "pnpm"
	execctl.Args = []string{"run"}
	execctl.Args = append(execctl.Args, cli.PnpmFlag.Values()...)

	if err := execctl.Exec(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	sig := <-sigch
	fmt.Printf("Received: %v\n", sig)
	if err := execctl.Kill(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

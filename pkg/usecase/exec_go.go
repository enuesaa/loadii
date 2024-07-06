package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/cli"
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func ExecGo(repos repository.Repos) {
	execctl := exec.New(repos)
	// TODO: 直感的でない
	execctl.Workdir = cli.GoFlag.Workdir()
	execctl.Command = "go"
	execctl.Args = []string{"run"}
	execctl.Args = append(execctl.Args, cli.GoFlag.Values()...)

	if err := execctl.Exec(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	// sig := <-sigCh
	// fmt.Printf("Received: %v\n", sig)
	// if err := execctl.Kill(); err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// }
}

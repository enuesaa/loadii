package usecase

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/enuesaa/tryserve/pkg/repository"
)

func RunGoApp(repos repository.Repos, path string) error {
	fmt.Printf("run: `go run %s`\n", path)

	cmd := exec.Command("go", "run", path)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Printf("running in pid %d\n", cmd.Process.Pid)

	if err := cmd.Wait(); err != nil {
		return err
	}
	fmt.Printf("stderr: %s\n", stderr.String())
	fmt.Printf("stdout: %s\n", stdout.String())

	return nil
}

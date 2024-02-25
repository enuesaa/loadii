package repository

import (
	"bytes"
	"fmt"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Exec(command string, args ...string) error
}
type CmdRepository struct{}

func (repo *CmdRepository) Exec(command string, args ...string) error {
	fmt.Printf("run: `%s %s`\n", command, args)

	cmd := exec.Command(command, args...)

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

package repository

import (
	"io"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Exec(workdir string, writer io.Writer, command string, args ...string) error
}
type CmdRepository struct{}

//TODO: refactor
func (repo *CmdRepository) Exec(workdir string, writer io.Writer, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = workdir

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

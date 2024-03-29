package repository

import (
	"io"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Exec(writer io.Writer, workdir string, command string, args []string) error
}
type CmdRepository struct{}

func (repo *CmdRepository) Exec(writer io.Writer, workdir string, command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = workdir

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

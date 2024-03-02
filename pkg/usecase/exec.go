package usecase

import (
	"fmt"
	"os/exec"

	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, commands []string) error {
	fmt.Printf("exec command: %v\n", commands)

	command := commands[0]
	args := commands[1:]

	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("output: %s\n", string(output))

	return nil
}

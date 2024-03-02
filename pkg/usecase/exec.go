package usecase

import (
	"fmt"
	"os/exec"
)

func Exec(commands []string) error {
	if len(commands) == 0 {
		return fmt.Errorf("please specify command")
	}
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

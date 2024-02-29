package usecase

import (
	"fmt"
	"os/exec"
)

func Exec(command string, args []string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("output: %s\n", string(output))

	return nil
}

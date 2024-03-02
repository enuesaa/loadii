package usecase

import (
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, commands []string) error {
	repos.Log.Info("exec command: %v", commands)

	command := commands[0]
	args := commands[1:]
	execctl := exec.New(repos)

	return execctl.Exec(command, args...)
}

package usecase

import (
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func Exec(repos repository.Repos, commands []string, workdir string) error {
	repos.Log.Info("Running %v", commands)

	command := commands[0]
	args := commands[1:]
	execctl := exec.New(repos)

	return execctl.Exec(command, args...)
}

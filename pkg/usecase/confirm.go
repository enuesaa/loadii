package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
)

func Confirm(repos repository.Repos, commands []string, autoApprove bool) error {
	repos.Log.Info("[PLAN] run the command `%v`", commands)

	if !autoApprove {
		answer, err := repos.Prompt.Confirm("Are you sure to run")
		if err != nil {
			return err
		}
		if !answer {
			return fmt.Errorf("not confirmed")
		}
	}

	return nil
}

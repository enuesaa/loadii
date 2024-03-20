package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/pkg/repository"
)

func Confirm(repos repository.Repos, commands []string, autoApprove bool) error {
	repos.Log.Info("[PLAN] run the command `%v`", commands)

	if !autoApprove {
		value, err := repos.Prompt.Ask("Are you sure to run (y/n)", "")
		if err != nil {
			return err
		}
		if value != "y" {
			return fmt.Errorf("not confirmed")
		}
	}

	return nil
}
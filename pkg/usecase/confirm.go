package usecase

import (
	"fmt"
	"strings"

	"github.com/enuesaa/loadii/pkg/repository"
)

type Plan struct {
	ServePath     string
	ServePort     int
	Commands      []string
	WatchIncludes []string
	WatchExcludes []string
}

func Confirm(repos repository.Repos, plan Plan, autoApprove bool) error {
	// print planning like fiber v2 message
	repos.Log.Info("┌────────────────────────────────────────────")
	repos.Log.Info("│ Loadii")
	repos.Log.Info("│")
	if plan.ServePath != "" {
		repos.Log.Info("│ Serve:   %s on %d", plan.ServePath, plan.ServePort)
	}
	if len(plan.Commands) > 0 {
		repos.Log.Info("│ Command: %s", strings.Join(plan.Commands, " "))
	}
	repos.Log.Info("│ WatchIncludes: %+v", plan.WatchIncludes)
	if len(plan.WatchExcludes) > 0 {
		repos.Log.Info("│ WatchExcludes: %+v", plan.WatchExcludes)
	}
	repos.Log.Info("└────────────────────────────────────────────")

	if !autoApprove {
		answer, err := repos.Log.Confirm("Are you sure to run")
		if err != nil {
			return err
		}
		if !answer {
			return fmt.Errorf("not confirmed")
		}
	}

	return nil
}

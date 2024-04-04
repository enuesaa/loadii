package usecase

import (
	"fmt"
	"strings"

	"github.com/enuesaa/loadii/pkg/repository"
)

type Plan struct {
	Workdir       string
	ServePath     string
	ServePort     int
	Commands      []string
	WatchIncludes []string
	WatchExcludes []string
}

func (p *Plan) Print(repos repository.Repos) {
	// print planning like fiber v2 message
	repos.Log.Info("┌────────────────────────────────────────────")
	repos.Log.Info("│ Loadii")
	repos.Log.Info("│")
	if p.ServePath != "" {
		repos.Log.Info("│ Serve:   %s on %d", p.ServePath, p.ServePort)
	}
	if len(p.Commands) > 0 {
		repos.Log.Info("│ Command: %s", strings.Join(p.Commands, " "))
	}
	repos.Log.Info("│ WatchIncludes: %+v", p.WatchIncludes)
	if len(p.WatchExcludes) > 0 {
		repos.Log.Info("│ WatchExcludes: %+v", p.WatchExcludes)
	}
	repos.Log.Info("└────────────────────────────────────────────")
}

func (p *Plan) Confirm(repos repository.Repos) error {
	answer, err := repos.Log.Confirm("Are you sure to run")
	if err != nil {
		return err
	}
	if !answer {
		return fmt.Errorf("not confirmed")
	}
	return nil
}

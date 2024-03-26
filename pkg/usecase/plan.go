package usecase

import "github.com/enuesaa/loadii/pkg/repository"

type PlanProps struct {
	ServePath string
	ServePort int
	Commands []string
	WatchIncludes []string
	WatchExcludes []string
}
func Plan(repos repository.Repos, props PlanProps) {
	// TODO: this is prototype
	// print planning like fiber v2 message
	repos.Log.Info("┌──────────────────────")
	repos.Log.Info("│ Loadii")
	repos.Log.Info("│")
	if props.ServePath != "" {
		repos.Log.Info("│ Serve:   %s on %d", props.ServePath, props.ServePort)
	} else {
		repos.Log.Info("│ Serve:   (not configured)")
	}
	if len(props.Commands) > 0 {
		repos.Log.Info("│ Command: %+v", props.Commands)
	} else {
		repos.Log.Info("│ Command: (not configured)")
	}
	repos.Log.Info("│ Watch:   includes: %+v, excludes: %+v", props.WatchIncludes, props.WatchExcludes)
	repos.Log.Info("└──────────────────────")
}

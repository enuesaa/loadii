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
	// print planning like fiber v2 message
	repos.Log.Info("┌──────────────────────")
	repos.Log.Info("│ Loadii")
	repos.Log.Info("│")
	if props.ServePath != "" {
		repos.Log.Info("│ Serve:   %s on %d", props.ServePath, props.ServePort)
	}
	if len(props.Commands) > 0 {
		repos.Log.Info("│ Command: %+v", props.Commands)
	}
	repos.Log.Info("│ Watch:   includes: %+v, excludes: %+v", props.WatchIncludes, props.WatchExcludes)
	repos.Log.Info("└──────────────────────")
}

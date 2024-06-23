package main

import (
	"fmt"
	"os"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
)

func main() {
	flags := parseArgs(os.Args)
	if flags.HasHelpFlag {
		fmt.Printf("%s\n", helpText)
		os.Exit(0)
	}
	if flags.HasVersionFlag {
		fmt.Printf("%s\n", versionText)
		os.Exit(0)
	}

	repos := repository.New()
	if flags.HasGoFlag {
		go func() {
			plan := usecase.Plan{
				Workdir:  flags.GoFlagPath,
				Commands: []string{"go", "run"},
			}
			if err := usecase.Exec(repos, plan); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		}()
	}
	if flags.HasPnpmFlag {
		go func() {
			plan := usecase.Plan{
				Workdir:  flags.PnpmFlagPath,
				Commands: []string{"pnpm", "run", flags.PnpmFlagScriptName},
			}
			if err := usecase.Exec(repos, plan); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		}()
	}

	if !flags.HasGoFlag && !flags.HasPnpmFlag {
		fmt.Printf("%s\n", helpText)
		os.Exit(0)
	}
}

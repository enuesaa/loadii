package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
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
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	if flags.HasGoFlag {
		go func() {
			execctl := exec.New(repos)
			execctl.Workdir = "."
			execctl.Command = "go"
			execctl.Args = []string{"run", flags.GoFlagPath}
			execctl.Args = append(execctl.Args, flags.GoArgs...)

			if err := execctl.Exec(); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			sig := <-sigCh
			fmt.Printf("Received: %v\n", sig)
			if err := execctl.Kill(); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		}()
	}
	if flags.HasPnpmFlag {
		go func() {
			execctl := exec.New(repos)
			execctl.Workdir = flags.PnpmFlagPath
			execctl.Command = "pnpm"
			execctl.Args = []string{"run", flags.PnpmFlagScriptName}

			if err := execctl.Exec(); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			sig := <-sigCh
			fmt.Printf("Received: %v\n", sig)
			if err := execctl.Kill(); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		}()
	}

	if flags.HasServeFlag {
		go func ()  {
			servectl := serve.New(repos)
			servectl.Port = 3000
			servectl.Basepath = "."

			fmt.Printf("Listening on %s", servectl.Addr())

			if err := servectl.Listen(); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			// sig := <-sigCh
			// fmt.Printf("Received: %v\n", sig)
			// if err := execctl.Kill(); err != nil {
			// 	fmt.Printf("Error: %s", err.Error())
			// }
		}()
	}

	if !flags.HasGoFlag && !flags.HasPnpmFlag {
		fmt.Printf("%s\n", helpText)
		os.Exit(0)
	}

	plan := usecase.Plan{
		WatchIncludes: []string{"."},
		WatchExcludes: []string{},
	}
	if err := usecase.Watch(repos, plan); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}

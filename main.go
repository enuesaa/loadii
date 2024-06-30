package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/loadii/pkg/cli"
	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/serve"
	"github.com/enuesaa/loadii/pkg/usecase"
)

func main() {
	cli.Parse(os.Args)

	if cli.HelpFlag.Has() {
		fmt.Printf("%s\n", cli.HelpText)
		os.Exit(0)
	}
	if cli.VersionFlag.Has() {
		fmt.Printf("%s\n", cli.VersionText)
		os.Exit(0)
	}

	repos := repository.New()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	if cli.GoFlag.Has() {
		go func() {
			execctl := exec.New(repos)
			execctl.Workdir = cli.GoFlag.Workdir()
			execctl.Command = "go"
			execctl.Args = []string{"run"}
			execctl.Args = append(execctl.Args, cli.GoFlag.Values()...)

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
	if cli.PnpmFlag.Has() {
		go func() {
			execctl := exec.New(repos)
			execctl.Workdir = cli.PnpmFlag.Workdir()
			execctl.Command = "pnpm"
			execctl.Args = []string{"run"}
			execctl.Args = append(execctl.Args, cli.PnpmFlag.Values()...)

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

	if cli.ServeFlag.Has() {
		go func ()  {
			servectl := serve.New(repos)
			servectl.Port = 3000
			servectl.Basepath = cli.ServeFlag.Workdir()

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

	if !cli.GoFlag.Has() && !cli.PnpmFlag.Has() {
		fmt.Printf("%s\n", cli.HelpText)
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

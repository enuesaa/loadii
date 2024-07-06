package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/loadii/pkg/cli"
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
)

func main() {
	cli.Parse(os.Args)

	if cli.HelpFlag.Has() {
		fmt.Printf("%s\n", cli.GetHelpText())
		os.Exit(0)
	}
	if cli.VersionFlag.Has() {
		fmt.Printf("%s\n", cli.GetVersionText())
		os.Exit(0)
	}

	repos := repository.New()
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM)

	if cli.GoFlag.Has() {
		go usecase.ExecGo(repos, sigch)
	}
	if cli.PnpmFlag.Has() {
		go usecase.ExecPnpm(repos, sigch)
	}

	if cli.ServeFlag.Has() {
		go usecase.Serve(repos, sigch)
	}

	if !cli.GoFlag.Has() && !cli.PnpmFlag.Has() {
		fmt.Printf("%s\n", cli.GetHelpText())
		os.Exit(0)
	}

	if err := usecase.Watch(repos, "."); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}

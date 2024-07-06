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
		go usecase.ExecGo(repos)
	}
	if cli.PnpmFlag.Has() {
		go usecase.ExecPnpm(repos)
	}

	if cli.ServeFlag.Has() {
		go usecase.Serve(repos)
	}

	if !cli.GoFlag.Has() && !cli.PnpmFlag.Has() {
		fmt.Printf("%s\n", cli.HelpText)
		os.Exit(0)
	}

	if err := usecase.Watch(repos, "."); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}

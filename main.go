package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/loadii/internal/cli"
	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/usecase"
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

	if err := usecase.Watch(repos, "."); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}

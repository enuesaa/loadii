package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

    "github.com/urfave/cli/v2"
	"github.com/enuesaa/loadii/internal/repository"
	"github.com/enuesaa/loadii/internal/usecase"
)

func main() {
	repos := repository.New()

	app := &cli.App{
        Name:  "loadii",
        Usage: "Instant web server",
        Action: func(*cli.Context) error {
			sigch := make(chan os.Signal, 1)
			signal.Notify(sigch, syscall.SIGTERM)		

			return usecase.Watch(repos, ".")
        },
    }

    if err := app.Run(os.Args); err != nil {
		log.Panicf("Error: %s", err.Error())
    }
}

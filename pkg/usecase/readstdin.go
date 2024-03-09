package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ReadStdin() error {
	signal.Ignore(syscall.SIGPIPE)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if !scanner.Scan() {
				break
			}
			text := scanner.Text()
			if text == "" {
				fmt.Printf("stdin empty\n")
				break
			}
			fmt.Printf("stdin: %s\n", text)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}()

	return nil
}

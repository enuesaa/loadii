package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ReadStdinAndPrintLoop() {
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

	signal.Ignore(syscall.SIGPIPE)
	for range 10 {
		time.Sleep(1 * time.Second)
		fmt.Printf("a %d\n", os.Getpid())
	}
}

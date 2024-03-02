package repository

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Info(format string, v ...any)
	Fatal(err error)
}
type LogRepository struct {}

func (repo *LogRepository) prefix() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

func (repo *LogRepository) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%s  %s\n", repo.prefix(), message)
}

func (repo *LogRepository) Fatal(err error) {
	log.Fatalf("%s  Error: %s\n", repo.prefix(), err.Error())
}

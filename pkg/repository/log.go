package repository

import (
	"log"
)


type LogRepositoryInterface interface {
	Info(format string, v ...any)
	Fatal(err error)
}
type LogRepository struct {}

func (repo *LogRepository) Info(format string, v ...any) {
	log.Printf(format, v...)
}

func (repo *LogRepository) Fatal(err error) {
	log.Fatalf("Error: %s", err.Error())
}

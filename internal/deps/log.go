package deps

import (
	"fmt"
	"log"
	"time"

	"github.com/erikgeiser/promptkit/textinput"
)

func init() {
	log.SetFlags(0)
}

type Log interface {
	Info(format string, v ...any)
	Fatal(err error)
	Ask(message string, defaultValue string) (string, error)
	Confirm(message string) (bool, error)
}
type LogImpl struct{}

func (repo *LogImpl) prefix() string {
	return time.Now().Local().Format("15:04:05")
}

func (repo *LogImpl) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%s  %s\n", repo.prefix(), message)
}

func (repo *LogImpl) Fatal(err error) {
	log.Fatalf("%s  Error: %s\n", repo.prefix(), err.Error())
}

func (repo *LogImpl) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue

	return input.RunPrompt()
}

func (repo *LogImpl) Confirm(message string) (bool, error) {
	message = fmt.Sprintf("%s  %s (y/n)", repo.prefix(), message)
	input := textinput.New(message)
	input.InitialValue = ""

	answer, err := input.RunPrompt()
	if err != nil {
		return false, err
	}

	return answer == "y", nil
}

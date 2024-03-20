package repository

import (
	"fmt"

	"github.com/erikgeiser/promptkit/textinput"
)

type PromptRepositoryInterface interface {
	Ask(message string, defaultValue string) (string, error)
	Confirm(message string) (bool, error)
}
type PromptRepository struct{}

func (prompt *PromptRepository) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue

	return input.RunPrompt()
}

func (prompt *PromptRepository) Confirm(message string) (bool, error) {
	message = fmt.Sprintf("%s (y/n)", message)
	input := textinput.New(message)
	input.InitialValue = ""

	answer, err := input.RunPrompt()
	if err != nil {
		return false, err
	}

	return answer == "y", nil
}

package repository

import (
	"github.com/erikgeiser/promptkit/textinput"
)

type PromptRepositoryInterface interface {
	Ask(message string, defaultValue string) (string, error)
}
type PromptRepository struct{}

func (prompt *PromptRepository) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue

	return input.RunPrompt()
}

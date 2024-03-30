package exec

import (
	"strings"

	"github.com/enuesaa/loadii/pkg/repository"
)

type CmdWriter struct {
	repos repository.Repos
}

// implements io.Writer
// TODO refactor. should stdout with colored output.
func (w *CmdWriter) Write(p []byte) (n int, err error) {
	text := string(p)
	texts := strings.Split(text, "\n")

	if texts[len(texts)-1] == "" {
		texts = texts[:len(texts)-1]
	}
	for _, line := range texts {
		w.repos.Log.Info(line)
	}

	return len(p), nil
}

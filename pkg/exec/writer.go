package exec

import (
	"strings"

	"github.com/enuesaa/loadii/pkg/repository"
)

type CmdWriter struct {
	repos repository.Repos
}

// implements io.Writer
func (w *CmdWriter) Write(p []byte) (n int, err error) {
	text := string(p)
	texts := strings.Split(text, "\n")

	for _, line := range texts {
		w.repos.Log.Info(line)
	}

	return len(p), nil
}

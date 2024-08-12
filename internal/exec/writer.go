package exec

import (
	"strings"
)

// implements io.Writer
// In the furture, this function should print colored output.
func (ctl *Execctl) Write(p []byte) (n int, err error) {
	text := string(p)
	lines := strings.Split(text, "\n")

	if len(lines) == 0 {
		return len(p), nil
	}

	lastI := len(lines) - 1
	if lines[lastI] == "" {
		lines = lines[:lastI]
	}
	for _, line := range lines {
		ctl.repos.Log.Info(line)
	}

	return len(p), nil
}

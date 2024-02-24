package usecase

import (
	"fmt"
	"strings"

	"github.com/enuesaa/tryserve/pkg/repository"
)

func RunApp(repos repository.Repos, path string) error {
	ext := repos.Fs.Ext(path)
	if ext == ".go" {
		return RunGoApp(repos, path)
	}

	data, err := repos.Fs.Read(path)
	if err != nil {
		return err
	}
	code := string(data)
	lines := strings.Split(code, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("file is empty")
	}

	firstline := lines[0]
	if !strings.HasPrefix(firstline, "#!") {
		return fmt.Errorf("failed to run file becuase this file does not contain shebang.")
	}

	return repos.Cmd.Exec(path)
}

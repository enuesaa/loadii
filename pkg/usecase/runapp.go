package usecase

import (
	"fmt"
	"slices"
	"strings"

	"github.com/enuesaa/loadii/pkg/exec"
	"github.com/enuesaa/loadii/pkg/repository"
)

func RunApp(repos repository.Repos, path string, args []string) error {
	execctl := exec.New(repos)

	ext := repos.Fs.Ext(path)
	if ext == ".go" {
		cmdargs := slices.Concat([]string{"run", path}, args)
		return execctl.Exec("go", cmdargs...)
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

	return execctl.Exec(path)
}

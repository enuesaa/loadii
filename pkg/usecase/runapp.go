package usecase

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/enuesaa/tryserve/pkg/repository"
)

func RunApp(repos repository.Repos, path string) error {
	ext := filepath.Ext(path)
	if ext == ".go" {
		return RunGoApp(repos, path)
	}

	data, err := os.ReadFile(path)
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

	cmd := exec.Command(path)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Printf("running in pid %d\n", cmd.Process.Pid)

	if err := cmd.Wait(); err != nil {
		return err
	}
	fmt.Printf("stderr: %s\n", stderr.String())
	fmt.Printf("stdout: %s\n", stdout.String())

	return nil
}

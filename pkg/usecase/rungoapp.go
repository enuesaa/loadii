package usecase

import (
	"github.com/enuesaa/tryup/pkg/repository"
)

func RunGoApp(repos repository.Repos, path string) error {
	return repos.Cmd.Exec("go", "run", path)
}

package usecase

import (
	"github.com/enuesaa/tryserve/pkg/repository"
)

func RunGoApp(repos repository.Repos, path string) error {
	return repos.Cmd.Exec("go", "run", path)
}

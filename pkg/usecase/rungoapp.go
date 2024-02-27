package usecase

import (
	"github.com/enuesaa/loadii/pkg/repository"
)

func RunGoApp(repos repository.Repos, path string) error {
	return repos.Cmd.Exec("go", "run", path)
}

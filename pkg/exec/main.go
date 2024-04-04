package exec

import (
	"github.com/enuesaa/loadii/pkg/repository"
)

func New(repos repository.Repos) Execctl {
	return Execctl{
		repos: repos,
		Workdir: "",
		Command: "",
		Args: []string{},
	}
}

type Execctl struct {
	repos repository.Repos
	Workdir string
	Command string
	Args []string
}

func (ctl *Execctl) Exec() error {
	return ctl.repos.Cmd.Exec(ctl, ctl.Workdir, ctl.Command, ctl.Args)
}

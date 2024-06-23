package exec

import (
	"os/exec"

	"github.com/enuesaa/loadii/pkg/repository"
)

func New(repos repository.Repos) Execctl {
	return Execctl{
		repos:   repos,
		Workdir: "",
		Command: "",
		Args:    []string{},
	}
}

type Execctl struct {
	repos   repository.Repos
	Workdir string
	Command string
	Args    []string
	Cmd     *exec.Cmd
}

func (ctl *Execctl) Exec() error {
	cmd, err := ctl.repos.Cmd.Exec(ctl, ctl.Workdir, ctl.Command, ctl.Args)
	if err != nil {
		return err
	}
	ctl.Cmd = cmd
	return nil
}

func (ctl *Execctl) Kill() error {
	if ctl.Cmd == nil {
		return nil
	}
	return ctl.repos.Cmd.Kill(ctl.Cmd)
}

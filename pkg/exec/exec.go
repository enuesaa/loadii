package exec

// TODO: refactor
func (ctl *Execctl) Exec(workdir string, command string, args ...string) error {
	writer := CmdWriter{
		repos: ctl.repos,
	}
	return ctl.repos.Cmd.Exec(workdir, &writer, command, args...)
}

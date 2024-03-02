package exec

func (ctl *Execctl) Exec(command string, args ...string) error {
	writer := CmdWriter {
		repos: ctl.repos,
	}
	return ctl.repos.Cmd.Exec(&writer, command, args...)
}

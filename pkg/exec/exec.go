package exec

func (ctl *Execctl) Exec(workdir string, commands []string) error {
	command := commands[0]
	args := commands[1:]

	writer := CmdWriter{
		repos: ctl.repos,
	}
	return ctl.repos.Cmd.Exec(&writer, workdir, command, args)
}

package repository

type Repos struct {
	Cmd CmdRepositoryInterface
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
	Prompt PromptRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsRepository{},
		Log: &LogRepository{},
		Prompt: &PromptRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsMockRepository{},
		Log: &LogRepository{},
		Prompt: &PromptRepository{},
	}
}

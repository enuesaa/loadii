package repository

type Repos struct {
	Cmd    CmdRepositoryInterface
	Fs     FsRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd:    &CmdRepository{},
		Fs:     &FsRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		Cmd:    &CmdRepository{},
		Fs:     &FsMockRepository{},
	}
}

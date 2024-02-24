package repository

type Repos struct {
	Fs     FsRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:     &FsRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		Fs:     &FsMockRepository{},
	}
}

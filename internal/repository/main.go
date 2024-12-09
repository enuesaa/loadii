package repository

import "go.uber.org/mock/gomock"

type Repos struct {
	Cmd CmdRepositoryInterface
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  &FsRepository{},
		Log: &LogRepository{},
	}
}

func NewMock(ctrl gomock.Controller) Repos {
	return Repos{
		Cmd: &CmdRepository{},
		Fs:  NewMockFsRepositoryInterface(&ctrl),
		Log: &LogRepository{},
	}
}

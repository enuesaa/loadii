package repository

import (
	"testing"

	"go.uber.org/mock/gomock"
)

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

func NewMock(t *testing.T) Repos {
	ctrl := gomock.NewController(t)

	return Repos{
		Cmd: &CmdRepository{},
		Fs:  NewMockFsRepositoryInterface(ctrl),
		Log: NewMockLogRepositoryInterface(ctrl),
	}
}

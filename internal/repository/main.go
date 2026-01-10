package repository

import (
	"testing"

	"go.uber.org/mock/gomock"
)

type Repos struct {
	Fs  FsRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:  &FsRepository{},
		Log: &LogRepository{},
	}
}

func NewMock(t *testing.T) Repos {
	ctrl := gomock.NewController(t)

	return Repos{
		Fs:  NewMockFsRepositoryInterface(ctrl),
		Log: NewMockLogRepositoryInterface(ctrl),
	}
}

package di

import (
	"testing"

	"go.uber.org/mock/gomock"
)

type Deps struct {
	Fs  Fs
	Log Log
}

func New() *Deps {
	return &Deps{
		Fs:  &FsImpl{},
		Log: &LogImpl{},
	}
}

func NewMock(t *testing.T) *Deps {
	ctrl := gomock.NewController(t)

	return &Deps{
		Fs:  NewMockFs(ctrl),
		Log: NewMockLog(ctrl),
	}
}

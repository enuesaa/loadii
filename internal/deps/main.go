package deps

import (
	"testing"

	"go.uber.org/mock/gomock"
)

type Repos struct {
	Fs  Fs
	Log Log
}

func New() *Repos {
	return &Repos{
		Fs:  &FsImpl{},
		Log: &LogImpl{},
	}
}

func NewMock(t *testing.T) *Repos {
	ctrl := gomock.NewController(t)

	return &Repos{
		Fs:  NewMockFs(ctrl),
		Log: NewMockLog(ctrl),
	}
}

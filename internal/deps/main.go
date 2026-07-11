package deps

import (
	"testing"

	"go.uber.org/mock/gomock"
)

type Container struct {
	Fs  Fs
	Log Log
}

func New() *Container {
	return &Container{
		Fs:  &FsImpl{},
		Log: &LogImpl{},
	}
}

func NewMock(t *testing.T) *Container {
	ctrl := gomock.NewController(t)

	return &Container{
		Fs:  NewMockFs(ctrl),
		Log: NewMockLog(ctrl),
	}
}

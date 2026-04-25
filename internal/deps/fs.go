package deps

import (
	"io"
	"os"
	"path/filepath"
)

type Fs interface {
	Ext(path string) string
	IsExist(path string) bool
	IsDir(path string) (bool, error)
	WorkDir() (string, error)
	Read(path string) ([]byte, error)
}
type FsImpl struct{}

func (repo *FsImpl) Ext(path string) string {
	return filepath.Ext(path)
}

func (repo *FsImpl) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsImpl) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsImpl) WorkDir() (string, error) {
	return os.Getwd()
}

func (repo *FsImpl) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

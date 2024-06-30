package cli

import (
	"fmt"
	"strings"
)

type Flag struct {
	Name string // like `-tag`
	Help string // help message
	MinValues int // minimum values count
	MaxValues int // maximum values count. if 0, this flag peforms bool flag.
	DefaultValues []string
	Workdir string // default value is `.`
	ReceiveWorkdir bool
}

func (f *Flag) Has() bool {
	for _, a := range Args {
		if a == f.Name {
			return true
		}
		if f.ReceiveWorkdir {
			if strings.HasPrefix(a, f.NameWithWorkdirPrefix()) {
				return true
			}
		}
	}
	return false
}

func (f *Flag) NameWithWorkdirPrefix() string {
	return fmt.Sprintf("%s:", f.Name)
}

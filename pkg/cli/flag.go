package cli

import (
	"fmt"
	"strings"
)

type Flag struct {
	Name string // like `-tag`
	Alias string // like `-t`
	Help string // help message
	MinValues int // minimum values count
	MaxValues int // maximum values count. if 0, this flag peforms bool flag.
	DefaultValues []string
	DefaultWorkdir string // default value is `.`
	ReceiveWorkdir bool
}

func (f *Flag) Has() bool {
	for _, a := range Args {
		if a == f.Name {
			return true
		}
		if f.HasAlias() {
			if a == f.Alias {
				return true
			}
		}
		if f.ReceiveWorkdir {
			if strings.HasPrefix(a, f.NameWithWorkdirPrefix()) {
				return true
			}
		}
	}
	return false
}

func (f *Flag) HasAlias() bool {
	return f.Alias != ""
}

func (f *Flag) Values() []string {
	list := make([]string, 0)

	useNext := false
	for _, a := range Args {
		if useNext {
			if strings.HasPrefix(a, "-") {
				return list
			}
			list = append(list, a)
		}
		if a == f.Name {
			useNext = true
			continue
		}
		if f.ReceiveWorkdir {
			if strings.HasPrefix(a, f.NameWithWorkdirPrefix()) {
				useNext = true
				continue
			}
		}
	}

	return list
}

func (f *Flag) Workdir() string {
	if !f.ReceiveWorkdir {
		return ""
	}
	workdir := f.DefaultWorkdir

	for _, a := range Args {
		if a == f.Name {
			return workdir
		}
		if f.ReceiveWorkdir {
			if strings.HasPrefix(a, f.NameWithWorkdirPrefix()) {
				return strings.ReplaceAll(a, f.NameWithWorkdirPrefix(), "")
			}
		}
	}

	return workdir
}

func (f *Flag) NameWithWorkdirPrefix() string {
	return fmt.Sprintf("%s:", f.Name)
}

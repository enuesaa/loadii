package cli

import (
	"strings"
)

type Flag struct {
	Name string // like `-tag`
	Alias string // like `-t`
	Help string // help message
	MinValues int // minimum values count
	MaxValues int // maximum values count. if 0, this flag peforms bool flag.
	DefaultValues []string
	DefaultWorkdir string // default value is `.` if "", workdir is not acceptable.
}

func (f *Flag) Has() bool {
	position := f.GetPosition()
	if position == -1 {
		return false
	}
	return len(Args) > position
}

func (f *Flag) GetPosition() int {
	for i, a := range Args {
		if a == f.Name {
			return i
		}
		if f.HasWorkdir() && strings.HasPrefix(a, f.Name + ":") {
			return i
		}
		if f.HasAlias() {
			if a == f.Alias {
				return i
			}
			if f.HasWorkdir() && strings.HasPrefix(a, f.Alias + ":") {
				return i
			}
		}
	}
	return -1
}

func (f *Flag) HasAlias() bool {
	return f.Alias != ""
}

func (f *Flag) HasWorkdir() bool {
	return f.DefaultWorkdir != ""
}

func (f *Flag) Values() []string {
	list := make([]string, 0)

	position := f.GetPosition()
	if position == -1 {
		return list
	}
	if len(Args) <= position {
		return list
	}
	
	rest := Args[position:]
	for _, a := range rest {
		if strings.HasPrefix(a, "-") {
			return list
		}
		list = append(list, a)
	}

	return list
}

func (f *Flag) Workdir() string {
	if f.DefaultWorkdir == "" {
		return ""
	}
	workdir := f.DefaultWorkdir

	position := f.GetPosition()
	if position == -1 {
		return workdir
	}
	if len(Args) <= position {
		return workdir
	}

	flag := Args[position]
	if flag == f.Name {
		return workdir
	}
	if f.HasAlias() && flag == f.Alias {
		return workdir
	}

	if strings.HasPrefix(flag, f.Name + ":") {
		return strings.ReplaceAll(flag, f.Name + ":", "")
	}
	if strings.HasPrefix(flag, f.Alias + ":") {
		return strings.ReplaceAll(flag, f.Alias + ":", "")
	}
	return workdir
}

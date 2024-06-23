package main

import (
	"strings"
)

var helpText = `A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags]

FLAGS:
  -go [path]               Run 'go run [path]'
  -pnpm [path]             Run 'pnpm run dev [path]'
  -pnpm:[script] [path]    Run 'pnpm run [script] [path]'
  -help                    Show help
  -version                 Print the version
`


type Flags struct {
	HasGoFlag bool
	GoFlagPath string
	HasPnpmFlag bool
	PnpmFlagScriptName string
	PnpmFlagPath string
	HasHelpFlag bool
	HasVersionFlag bool
}
func parseArgs(args []string) Flags {
	flags := Flags{
		HasGoFlag: false,
		GoFlagPath: ".",
		HasPnpmFlag: false,
		PnpmFlagScriptName: "dev",
		PnpmFlagPath: ".",
		HasHelpFlag: false,
		HasVersionFlag: false,
	}

	for _, arg := range args[1:] {
		if arg == "-help" {
			flags.HasHelpFlag = true
			continue
		}

		if arg == "-version" {
			flags.HasVersionFlag = true
			continue
		}
	}

	checkNextAsGoFlagPath := false
	for _, arg := range args[1:] {
		if checkNextAsGoFlagPath {
			if strings.HasPrefix(arg, "-") {
				break
			}
			flags.GoFlagPath = arg
			break
		}
		if arg == "-go" {
			flags.HasGoFlag = true
			checkNextAsGoFlagPath = true
			continue
		}
	}

	checkNextAsPnpmFlagPath := false
	for _, arg := range args[1:] {
		if checkNextAsPnpmFlagPath {
			if strings.HasPrefix(arg, "-") {
				break
			}
			flags.PnpmFlagPath = arg
			break
		}
		if strings.HasPrefix(arg, "-pnpm:") {
			flags.HasPnpmFlag = true
			flags.PnpmFlagScriptName = strings.ReplaceAll(arg, "-pnpm:", "")
			checkNextAsPnpmFlagPath = true
			continue
		}
		if strings.HasPrefix(arg, "-pnpm") {
			flags.HasPnpmFlag = true
			checkNextAsPnpmFlagPath = true
			continue
		}
	}

	return flags
}

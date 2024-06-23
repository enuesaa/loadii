package main

import (
	"strings"
)

var helpText = `A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags]

FLAGS:
  -go:[path] [args]        Run 'go run [path] [args]'
  -pnpm:[path]             Run 'pnpm run dev [path]'
  -pnpm:[path] [script]    Run 'pnpm run [script] [path]'
  -help                    Show help
  -version                 Print the version`

var versionText = `loadii version 0.0.7`


type Flags struct {
	HasGoFlag bool
	GoFlagPath string
	GoArgs string
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
		GoArgs: "",
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
			flags.GoArgs = arg
			break
		}
		if strings.HasPrefix(arg, "-go:") {
			flags.HasGoFlag = true
			flags.GoFlagPath = strings.ReplaceAll(arg, "-go:", "")
			continue
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
			flags.PnpmFlagScriptName = arg
			break
		}
		if strings.HasPrefix(arg, "-pnpm:") {
			flags.HasPnpmFlag = true
			flags.PnpmFlagPath = strings.ReplaceAll(arg, "-pnpm:", "")
			checkNextAsPnpmFlagPath = true
			continue
		}
		if arg == "-pnpm" {
			flags.HasPnpmFlag = true
			checkNextAsPnpmFlagPath = true
			continue
		}
	}

	return flags
}

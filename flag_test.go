package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	cases := []struct{
		Args []string
		Expected Flags
	}{
		{
			Args: []string{"loadii", "-go"},
			Expected: Flags{
				HasGoFlag: true,
				GoFlagPath: ".",
				GoArgs: []string{},
				HasPnpmFlag: false,
				PnpmFlagScriptName: "dev",
				PnpmFlagPath: ".",
				HasHelpFlag: false,
				HasVersionFlag: false,
			},
		},
		{
			Args: []string{"loadii", "-go", "sub"},
			Expected: Flags{
				HasGoFlag: true,
				GoFlagPath: ".",
				GoArgs: []string{"sub"},
				HasPnpmFlag: false,
				PnpmFlagScriptName: "dev",
				PnpmFlagPath: ".",
				HasHelpFlag: false,
				HasVersionFlag: false,
			},
		},
		{
			Args: []string{"loadii", "-go", "-pnpm"},
			Expected: Flags{
				HasGoFlag: true,
				GoFlagPath: ".",
				GoArgs: []string{},
				HasPnpmFlag: true,
				PnpmFlagScriptName: "dev",
				PnpmFlagPath: ".",
				HasHelpFlag: false,
				HasVersionFlag: false,
			},
		},
		{
			Args: []string{"loadii", "-help"},
			Expected: Flags{
				HasGoFlag: false,
				GoFlagPath: ".",
				GoArgs: []string{},
				HasPnpmFlag: false,
				PnpmFlagScriptName: "dev",
				PnpmFlagPath: ".",
				HasHelpFlag: true,
				HasVersionFlag: false,
			},
		},
		{
			Args: []string{"loadii", "-version"},
			Expected: Flags{
				HasGoFlag: false,
				GoFlagPath: ".",
				GoArgs: []string{},
				HasPnpmFlag: false,
				PnpmFlagScriptName: "dev",
				PnpmFlagPath: ".",
				HasHelpFlag: false,
				HasVersionFlag: true,
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, parseArgs(c.Args), c.Expected)
	}
}


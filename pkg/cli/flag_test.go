package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	Parse([]string{"loadii", "-go"})
	assert.Equal(t, GoFlag.Has(), true)

	Parse([]string{"loadii", "-go", "sub"})
	assert.Equal(t, GoFlag.Values(), []string{"sub"})

	Parse([]string{"loadii", "-go", "-pnpm"})
	assert.Equal(t, GoFlag.Has(), true)

	Parse([]string{"loadii", "-go", "-pnpm"})
	assert.Equal(t, PnpmFlag.Has(), true)

	Parse([]string{"loadii", "-go", "-pnpm"})
	assert.Equal(t, ServeFlag.Has(), false)

	Parse([]string{"loadii", "-go", "-help"})
	assert.Equal(t, HelpFlag.Has(), true)

	Parse([]string{"loadii", "-go", "-version"})
	assert.Equal(t, VersionFlag.Has(), true)
}


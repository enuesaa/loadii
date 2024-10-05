package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	Parse([]string{"loadii", "-go", "-help"})
	assert.Equal(t, HelpFlag.Has(), true)

	Parse([]string{"loadii", "-go", "-version"})
	assert.Equal(t, VersionFlag.Has(), true)

	Parse([]string{"loadii", "-go", "-v"})
	assert.Equal(t, VersionFlag.Has(), true)
}


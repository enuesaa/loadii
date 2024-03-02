package serve

import (
	"testing"

	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestConvertath(t *testing.T) {
	servectl := New(repository.New())

	assert.Equal(t, "index.html", servectl.convertPath("/"))
	assert.Equal(t, "index.html", servectl.convertPath("/index.html"))
	assert.Equal(t, "aa/index.html", servectl.convertPath("/aa/"))
}

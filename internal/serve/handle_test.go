package serve

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/loadii/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TODO: repository.NewMock() については要検討。いちいちキャストしたくない
func TestHandleMainRoute(t *testing.T) {
	repos := repository.NewMock(t)

	indexhtml := []byte("<html><body>hello</body></html>")
	repos.Fs.(*repository.MockFsRepositoryInterface).EXPECT().Read("index.html").Return(indexhtml, nil)
	repos.Log.(*repository.MockLogRepositoryInterface).EXPECT().Info(gomock.Any(), gomock.Any())

	servectl := New(repos)
	app := servectl.App()

    req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	res, err := app.Test(req)
	require.NoError(t, err)

	resbody, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	assert.Equal(t, indexhtml, resbody)
}

func TestConvertPath(t *testing.T) {
	servectl := New(repository.New())

	assert.Equal(t, "index.html", servectl.convertPath("/"))
	assert.Equal(t, "index.html", servectl.convertPath("/index.html"))
	assert.Equal(t, "aa/index.html", servectl.convertPath("/aa/"))
}

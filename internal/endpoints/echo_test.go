package endpoints

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/rest/web"
)

func TestPostEcho(t *testing.T) {
	ws := web.DefaultService()

	PostEcho(ws)
	ts := httptest.NewServer(ws)
	defer ts.Close()

	r := strings.NewReader(`{"request":"hello world"}`)
	res, err := http.Post(ts.URL+"/echo", "application/json", r)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t,
		`{"result":"hello world"}`+"\n",
		string(data),
	)
}

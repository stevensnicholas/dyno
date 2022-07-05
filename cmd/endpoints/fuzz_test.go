package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/rest/web"
)

func TestPostFuzz(t *testing.T) {
	ws := web.DefaultService()
	
	OpenAPIJSON, err := os.ReadFile("../../demo_server/swagger.json")
	assert.Nil(t, err)
	
	PostFuzz(ws)
	ts := httptest.NewServer(ws)
	defer ts.Close()

	requestText := fmt.Sprintf("request:%s", OpenAPIJSON)
	request, err := json.Marshal(requestText)
	r := strings.NewReader(request)
	res, err := http.Post(ts.URL+"/fuzz_client", "application/json", r)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)
	result := fmt.Sprintf(`{"result":%s`, OpenAPIJSON)
	assert.Equal(t, 
		result, 
		string(data),
	)
}
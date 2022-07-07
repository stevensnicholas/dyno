package endpoints

import (
	"context"
	"io"
	"os"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/rest/web"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSSendMessageImpl struct{}

func (dt SQSSendMessageImpl) GetQueueUrl(ctx context.Context,
	params *sqs.GetQueueUrlInput,
	optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error) {

	// URLs look like:
	//    https://sqs.REGION.amazonaws.com/ACCOUNT#/QUEUE-NAME
	prefix := "https://sqs.REGION.amazonaws.com/ACCOUNT#/"

	output := &sqs.GetQueueUrlOutput{
		QueueUrl: aws.String(prefix + "aws-docs-example-queue-url1"),
	}

	return output, nil
}

func (dt SQSSendMessageImpl) SendMessage(ctx context.Context,
	params *sqs.SendMessageInput,
	optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {

	output := &sqs.SendMessageOutput{
		MessageId: aws.String("aws-docs-example-messageID"),
	}

	return output, nil
}

type Config struct {
	QueueName string `json:"QueueName"`
}

var configFileName = "config.json"

var globalConfig Config


func TestPostFuzzValid(t *testing.T) {
	ws := web.DefaultService()
	
	OpenAPIJSON, err := os.ReadFile("../../demo_server/swagger.json")
	assert.Nil(t, err)
	
	PostFuzz(ws)
	ts := httptest.NewServer(ws)
	defer ts.Close()
	request := string(OpenAPIJSON)
	
	r := strings.NewReader(request)
	res, err := http.Post(ts.URL+"/fuzz_client", "application/json", r)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)

	result, err := os.ReadFile("swagger_test.txt")
	assert.Nil(t, err)

	assert.Equal(t, 
		string(result), 
		string(data),
	)
}

func TestPostFuzzInvalid(t *testing.T) {
	ws := web.DefaultService()
	
	PostFuzz(ws)
	ts := httptest.NewServer(ws)
	defer ts.Close()

	res, err := http.Post(ts.URL+"/fuzz_client", "application/json", strings.NewReader(""))
	assert.Nil(t, err)
	assert.Equal(t, 400, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, 
		`{"status":"INVALID_ARGUMENT","error":"invalid argument: missing request body"}` + "\n", 
		string(data),
	)
}

func TestPostFuzzNotOpenAPIJSON(t *testing.T) {
	ws := web.DefaultService()

	PostFuzz(ws)
	ts := httptest.NewServer(ws)
	defer ts.Close()

	r := strings.NewReader(`{"request":"hello world"}`)
	res, err := http.Post(ts.URL+"/echo", "application/json", r)
	assert.Nil(t, err)
	assert.Equal(t, 400, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t,
		`{"error":"not a valid request"}`+"\n",
		string(data),
	)
}
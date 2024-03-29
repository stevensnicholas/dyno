package main

import (
	"archive/zip"
	"bytes"
	"context"
	"dyno/internal/issue"
	"dyno/internal/logger"
	"dyno/internal/parse"
	"dyno/internal/platform"
	"dyno/internal/result"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/go-github/v45/github"
)

func main() {
	lambda.Start(handler)
}

var (
	s3session *s3.S3
)

const (
	REGION = "ap-southeast-2"
)

type Message struct {
	Key        *string `json:"key,omitempty"`
	BucketName *string `json:"bucketName,omitempty"`
	UUID       *string `json:"uuid,omitempty"`
	Token      *string `json:"token,omitempty"`
	Owner      *string `json:"owner,omitempty"`
	Repo       *string `json:"repo,omitempty"`
}

type SQSMessage struct {
	Type      *string `json:"Type,omitempty"`
	MessageID *string `json:"MessageId,omitempty"`
	TopicArn  *string `json:"TopicArn,omitempty"`
	Message   *string `json:"Message,omitempty"`
	Timestamp *string `Timestamp:"Type,omitempty"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	logLevel := "info"
	log, err := logger.ConfigureDevelopmentLogger(logLevel)
	var sqsMessage SQSMessage
	var message Message
	if err != nil {
		panic(err)
	}
	defer log.Sync()
	logger.Info("Function invoked!")
	if len(sqsEvent.Records) == 0 {
		logger.Warn("No SQS events")
	}
	for _, sqsEvent := range sqsEvent.Records {
		err = json.Unmarshal([]byte(sqsEvent.Body), &sqsMessage)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(*sqsMessage.Message), &message)
		if err != nil {
			panic(err)
		}
	}
	s3Key := *message.Key

	obj := getObject(&s3Key, message.BucketName)

	body, err := ioutil.ReadAll(obj.Body)

	if err != nil {
		panic(err)
	}

	reader, err := zip.NewReader(bytes.NewReader(body), *obj.ContentLength)
	if err != nil {
		panic(err)
	}

	results := readZipFile(reader)
	issues := issue.CreateIssues("RESTler", results)
	githubClient := platform.CreateGithubClient(ctx, message.Token)

	for i := 0; i < len(issues); i++ {
		_, _, err := githubClient.Issues.Create(ctx, *message.Owner, *message.Repo, &github.IssueRequest{
			Title:     issues[i].Title,
			Body:      platform.FormatFuzzBody(&issues[i]),
			Labels:    issues[i].Labels,
			Assignee:  issues[i].Assignee,
			State:     issues[i].State,
			Milestone: issues[i].Milestone,
		})

		if err != nil {
			panic(err)
		}
	}

	if len(issues) != 0 {
		logger.Info("Issues sent to client")
	} else {
		logger.Warn("No issues found in API")
	}

	return nil
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

// getObject uses the s3 key and bucketName to locate the object
// where the fuzzer results are located allowing for raw result parsing
// Inputted is the key and bucketName and returned is the s3 object
func getObject(key *string, bucketName *string) *s3.GetObjectOutput {
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(*bucketName),
		Key:    aws.String(*key),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

// readZipFile reads the s3 object results.zip and decompresses it.
// Reading the files in memory and using the parser to turn fuzz results into
// dynoResults that are then used to create real issues.
// Inputted is the zip.Reader and returned is a list of results
func readZipFile(reader *zip.Reader) []result.DynoResult {
	results := []result.DynoResult{}
	for _, f := range reader.File {
		read, err := f.Open()
		if err != nil {
			panic(err)
		}
		bugFileName := strings.Split(f.Name, "/")
		fuzzError := strings.Split(bugFileName[len(bugFileName)-1], "_")
		if len(fuzzError) > 1 {
			fileContents, _ := ioutil.ReadAll(read)
			rawResults := parse.ParseRestlerFuzzResults(string(fileContents), fuzzError)
			for i := 0; i < len(rawResults); i++ {
				results = append(results, rawResults[i])
			}
		}
	}
	return results
}

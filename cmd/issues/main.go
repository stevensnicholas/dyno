package main

import (
	"dyno/internal/result"
	"bytes"
	"archive/zip"
	"context"
	"dyno/internal/issue"
	"dyno/internal/logger"
	"dyno/internal/parse"
	"dyno/internal/platform"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/go-github/v45/github"
	"io/ioutil"
	"strings"
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
	Location   *string `json:"location,omitempty"`
	BucketName *string `json:"bucketName,omitempty"`
	UUID       *string `json:"uuid,omitempty"`
	Token      *string `json:"token,omitempty"`
	Owner      *string `json:"owner,omitempty"`
	Repo       *string `json:"repo,omitempty"`
}

type SQSMessage struct {
	Type   *string `json:"Type,omitempty"`
	MessageID *string `json:"MessageId,omitempty"`
	TopicArn   *string `json:"TopicArn,omitempty"`
	Message *string `json:"Message,omitempty"`
	Timestamp   *string `Timestamp:"Type,omitempty"`
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
		logger.Info("No SQS events")
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
	filename := *message.Location


	obj := getObject(&filename, message.BucketName)

	body, err := ioutil.ReadAll(obj.Body)

	if err != nil {
		panic(err)
	}
	
	reader, err := zip.NewReader(bytes.NewReader(body), *obj.ContentLength)
	if err != nil {
		panic(err)
	}
	results := []result.DynoResult{}
	for _, f := range reader.File {
		read, err := f.Open()
		if err != nil {
			panic(err)
		}
		bugFileName := strings.Split(f.Name, "/")
		if len(bugFileName) > 0 {
			fuzzError := strings.Split(bugFileName[1], "_")
			if len(fuzzError) > 1 {
				fileContents, _ := ioutil.ReadAll(read)
				rawResults := parse.ParseRestlerFuzzResults(string(fileContents), fuzzError)
				for _, result := range rawResults {
					results = append(results, result)
				}
			}
		}
	}

	issues := issue.CreateIssues(results)
	githubClient := platform.CreateGithubClient(ctx, message.Token)

	for _, issue := range issues {
		githubClient.Issues.Create(ctx, *message.Owner, *message.Repo, &github.IssueRequest{
			Title:     issue.Title,
			Body:      platform.FormatFuzzBody(&issue),
			Labels:    issue.Labels,
			Assignee:  issue.Assignee,
			State:     issue.State,
			Milestone: issue.Milestone,
		})
	}

	if len(issues) != 0 {
		logger.Info("Issues sent to client")
	} else {
		logger.Info("No issues found in API")
	}

	return nil
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func getObject(filename *string, bucketName *string) *s3.GetObjectOutput {
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(*bucketName),
		Key:    aws.String(*filename),
	})

	if err != nil {
		panic(err)
	}

	return resp
}
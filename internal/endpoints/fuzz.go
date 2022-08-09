package endpoints

import (
	"context"
	"dyno/internal/logger"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"os"
	"strings"
)

type cliInput struct {
	Data []byte `json:"result"`
}

type cliOutput struct {
	Result string `json:"result"`
}

func Fuzz(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*cliInput)
		var out = output.(*cliOutput)

		logger.Infof("Received client OpenAPI File")
		out.Result = "Saved"
		bod := strings.NewReader(string(in.Data))

		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-2")}))
		uploader := s3manager.NewUploader(sess)

		u := uuid.New()
		key := fmt.Sprintf("Open-Api-Files/%s", u.String())
		bucketName := os.Getenv("open_api_s3_name")

		logger.Infof("bucket is %s", bucketName)
		_, ierr := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
			Body:   bod,
		})

		s3URI := fmt.Sprintf("{\"s3_location\":\"s3://%s/%s\",\"uuid\":\"%s\"}", bucketName, key, u.String())

		if ierr != nil {
			logger.Infof("There was an issue uploading to s3: %s", ierr.Error())
		}

		//add sqs message
		sess2 := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		svc := sqs.New(sess2)

		qURL := os.Getenv("open_api_sqs_url")

		result, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  aws.String(s3URI),
			QueueUrl:     &qURL,
		})

		if err != nil {
			logger.Infof("There was an issue adding event to sqs: %s", err.Error())
		}

		logger.Infof("Success", *result.MessageId)

		return nil
	}

	u := usecase.NewIOI(new(cliInput), new(cliOutput), handler)
	u.SetTitle("Open Api Fuzz")
	u.SetDescription("Recieves the open-api file from client and adds to s3")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/openapi", u)
}

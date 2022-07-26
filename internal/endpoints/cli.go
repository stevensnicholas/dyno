package endpoints

import (
	"context"
	"dyno/internal/logger"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type cliInput struct {
	HTTPMethod      string            `json:"httpMethod"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

type cliOutput struct {
	Result string `json:"result"`
}

func recieveFile(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*cliInput)
		var out = output.(*cliOutput)

		logger.Infof("Received client OpenAPI FIle: %s", in.Body)
		out.Result = in.Body
		bod := strings.NewReader(in.Body)

		sess := session.Must(session.NewSession())
		uploader := s3manager.NewUploader(sess)

		//will replace with the openapi bucket when created
		u := uuid.New()
		key := fmt.Sprintf("File/%s", u.String())
		_, ierr := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("test-store-swagger"),
			Key:    aws.String("clientOpenApiFie"),
			Body:   bod,
		})

		if ierr != nil {
			logger.Infof("There was an issue uploading to s3: %s", ierr.Error())
		}

		return nil
	}

	u := usecase.NewIOI(new(cliInput), new(cliOutput), handler)
	u.SetTitle("OpenApiFile")
	u.SetDescription("Recieves the open-api file from client and adds to s3")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/cli", u)
}

package endpoints

import (
	"context"
	"dyno/internal/logger"
	"strings"
  "fmt"
	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
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
			Region: aws.String("ap-southeast-2")},))
		uploader := s3manager.NewUploader(sess)

		u := uuid.New()
		key := fmt.Sprintf("Open-Api-Files/%s", u.String())
		_, ierr := uploader.Upload(&s3manager.UploadInput{
			//will use the bucket name as variable
			Bucket: aws.String("test-store-swagger"),
			Key:    aws.String(key),
			Body:   bod,
		})

		if ierr != nil {
			logger.Infof("There was an issue uploading to s3: %s", ierr.Error())
		}

		return nil
	}

	u := usecase.NewIOI(new(cliInput), new(cliOutput), handler)
	u.SetTitle("Open Api Fuzz")
	u.SetDescription("Recieves the open-api file from client and adds to s3")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/openapi", u)
}
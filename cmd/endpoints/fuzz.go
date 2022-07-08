package endpoints

import (
	"context"
	"golambda/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSSendMessageAPI defines the interface for the GetQueueUrl and SendMessage functions.
// We use this interface to test the functions using a mocked service.
type SQSSendMessageAPI interface {
	GetQueueUrl(ctx context.Context,
		params *sqs.GetQueueUrlInput,
		optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)

	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

// Functions for SQS AWS service to send message

// GetQueueURL gets the URL of an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a GetQueueUrlOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to GetQueueUrl.
func GetQueueURL(c context.Context, api SQSSendMessageAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

// SendMsg sends a message to an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.
func SendMsg(c context.Context, api SQSSendMessageAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(c, input)
}

// Receiving a OpenAPI json file 
// TODO check if this should be the text of the swagger file or an upload of the swagger file
type PostFuzzInput struct {
	OpenAPIVersion string `json:"openapi"`
	Information map[string]interface{} `json:"info"`
	Servers interface{} `json:"servers"`
	Paths map[string]interface{} `json:"paths"`
	Components map[string]interface{} `json:"components"`
}
// Returning the OpenAPI json file
// TODO discuss if this should actually be the fuzzing results
type PostFuzzOutput struct {
	OpenAPIVersion string `json:"openapi"`
	Information map[string]interface{} `json:"info"`
	Servers interface{} `json:"servers"`
	Paths map[string]interface{} `json:"paths"`
	Components map[string]interface{} `json:"components"`
}

// Post Fuzz Creates a new endpoint at /fuzz_client 
// That is utilized to obtain a repo's openapi json file 
// This is then added to a sqs to be used by Restler 
// Inputs:
// 		param service Web server 
// Outputs:
// 		returns nil nothing 
func PostFuzz(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*PostFuzzInput)
		var out = output.(*PostFuzzOutput)

		// TODO Error check if the recieved json file

		logger.Infof("Received OpenAPI JSON request: Version of OpenAPI %s", in.OpenAPIVersion)
		// Returning the json openapi file that has been received
		out.OpenAPIVersion = in.OpenAPIVersion
		out.Information = in.Information
		out.Servers = in.Servers
		out.Paths = in.Paths
		out.Components = in.Components

		return nil
	}

	// OpenAPI Documentation
	u := usecase.NewIOI(new(PostFuzzInput), new(PostFuzzOutput), handler)
	u.SetTitle("Fuzz_API")
	u.SetExpectedErrors(status.InvalidArgument, status.Aborted)

	service.Post("/fuzz_client", u)

}
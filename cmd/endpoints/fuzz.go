package endpoints

import (
	"context"
	"golambda/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)
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

// Function Post Fuzz 
// Creates a new endpoint at /fuzz_client 
// That is utilized to obtain a repo's openapi json file 
// This is then added to a sqs to be used by Restler 
// @param service Web server 
// returns nil nothing 
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
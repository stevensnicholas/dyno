package endpoints

import (
	"context"
	"golambda/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type PostFuzzInput struct {
	Request string `json:"request"`
}

type PostFuzzOutput struct {
	Result string `json:"result"`
}

func PostFuzz(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*PostFuzzInput)
		var out = output.(*PostFuzzOutput)

		logger.Infof("Received OpenAPI JSON request: %.5s")
		out.Result = in.Request // Post should return the whole openapi blueprint

		return nil
	}

	// OpenAPI Documentation
	u := usecase.NewIOI(new(PostFuzzInput), new(PostFuzzOutput), handler)
	u.SetTitle("Fuzz_API")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/fuzz_client", u)

}
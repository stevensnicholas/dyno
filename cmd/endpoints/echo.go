package endpoints

import (
	"context"
	"golambda/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type PostEchoInput struct {
	Request string `json:"request"`
}

type PostEchoOutput struct {
	Result string `json:"result"`
}

func PostEcho(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*PostEchoInput)
		var out = output.(*PostEchoOutput)

		logger.Infof("Received echo request: %s", in.Request)
		out.Result = in.Request

		return nil
	}

	u := usecase.NewIOI(new(PostEchoInput), new(PostEchoOutput), handler)

	u.SetTitle("Echo")
	u.SetDescription("Returns the same string as provided")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/echo", u)
}

package endpoints

import (
	"context"
	"dyno/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetFuzzBugsInput struct {
	Request string `json:"request"`
}

type GetFuzzBugsOutput struct {
	Fuzzes []Fuzz `json:"fuzzes"`
}

type FuzzBugs struct {
	Title     string `json:"title"`
	ErrorType string `json:"errorType"`
	EndPoints string `json:"EndPoints"`
}

func GetFuzzesBug(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*GetFuzzBugsInput)
		var out = output.(*GetFuzzBugsOutput)

		logger.Infof("Received echo request: %s", in.Request)
		out.Fuzzes = []Fuzz{}

		return nil
	}

	u := usecase.NewIOI(new(GetFuzzesInput), new(GetFuzzesOutput), handler)

	u.SetTitle("Get Fuzz Bugs")
	u.SetDescription("Returns the list of all the Fuzz Bugs")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/fuzzBugs", u)
}

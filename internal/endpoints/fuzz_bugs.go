package endpoints

import (
	"context"
	"dyno/internal/logger"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetFuzzBugsInput struct {
	UUID string `json:"uuid "`
}

type GetFuzzBugsOutput struct {
	Fuzzes []Fuzzb `json:"fuzzes"`
}

type Fuzzb struct {
	Title     string `json:"title"`
	ErrorType string `json:"errorType"`
	EndPoints string `json:"endPoints"`
}

func GetFuzzesBug(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*GetFuzzBugsInput)
		var out = output.(*GetFuzzBugsOutput)

		logger.Infof("Received echo request: %s", in.UUID)
		out.Fuzzes = []Fuzzb{}

		return nil
	}

	u := usecase.NewIOI(new(GetFuzzBugsInput), new(GetFuzzBugsOutput), handler)

	u.SetTitle("Get Fuzz Bugs")
	u.SetDescription("Returns the list of all the Fuzz Bugs")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/fuzzBugs", u)
}

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
	Fuzzes []Fuzzbugs `json:"fuzzes"`
}

type Fuzzbugs struct {
	Title     string `json:"title"`
	ErrorType string `json:"errorType"`
	Endpoints string `json:"endpoints"`
}

func GetFuzzesBug(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*GetFuzzBugsInput)
		var out = output.(*GetFuzzBugsOutput)

		logger.Infof("Received echo request: %s", in.UUID)
		out.Fuzzes = []Fuzzbugs{}

		return nil
	}

	u := usecase.NewIOI(new(GetFuzzBugsInput), new(GetFuzzBugsOutput), handler)

	u.SetTitle("Get Fuzz Bugs")
	u.SetDescription("Returns the list of all the Fuzz Bugs")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/fuzzBugs", u)
}

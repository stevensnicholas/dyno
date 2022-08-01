package endpoints

import (
	"context"
	"dyno/internal/logger"
	"time"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetFuzzesInput struct {
	Request string `json:"request"`
}

type GetFuzzesOutput struct {
	Fuzzes []Fuzz `json:"fuzzes"`
}

type Fuzz struct {
	Id       string    `json:"id"`
	Date     time.Time `json:"time"`
	BugCount uint      `json:"bugCount"`
}

func GetFuzzes(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*GetFuzzesInput)
		var out = output.(*GetFuzzesOutput)

		logger.Infof("Received echo request: %s", in.Request)
		out.Fuzzes = []Fuzz{}

		return nil
	}

	u := usecase.NewIOI(new(GetFuzzesInput), new(GetFuzzesOutput), handler)

	u.SetTitle("Get Fuzzes")
	u.SetDescription("Returns the list of all the times the fuzzer ran")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/fuzzes", u)
}

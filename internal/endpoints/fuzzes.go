package endpoints

import (
	"context"
	"time"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetFuzzesInput struct {
	Request string `json:"request"`
}

type GetFuzzesOutput struct {
	Fuzzes []Fuzzes `json:"fuzzes"`
}

type Fuzzes struct {
	ID       string    `json:"id"`
	Date     time.Time `json:"time"`
	BugCount uint      `json:"bugCount"`
}

func GetFuzzes(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var out = output.(*GetFuzzesOutput)
		out.Fuzzes = []Fuzzes{}

		return nil
	}

	u := usecase.NewIOI(new(GetFuzzesInput), new(GetFuzzesOutput), handler)

	u.SetTitle("Get Fuzzes")
	u.SetDescription("Returns the list of all the times the fuzzer ran")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/fuzzes", u)
}

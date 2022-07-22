package endpoints

import (
	"context"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type GetCodeInput struct {
	Code string `query:"code"`
}

type GetCodeOutput struct {
	Result string `json:"result"`
}

func GetCode(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var out = output.(*GetCodeOutput)

		out.Result = "code"

		return nil
	}

	u := usecase.NewIOI(new(GetCodeInput), new(GetCodeOutput), handler)

	u.SetTitle("code")
	u.SetDescription("Return code")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/login", u)
}

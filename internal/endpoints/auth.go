package endpoints

import (
	"context"
	"dyno/internal/authentication"
	"dyno/internal/logger"
	"fmt"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type AuthInput struct {
	Code string `query:"code"`
}

type AuthOutput struct {
	Result string `json:"token"`
}

func Authentication(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*AuthInput)
		var out = output.(*AuthOutput)
		fmt.Println(in)
		var err error
		var code = in.Code
		fmt.Println("code", code)
		var tokenAuthURL = authentication.GetTokenAuthURL(code)
		var token *authentication.Token
		if token, err = authentication.GetToken(tokenAuthURL); err != nil {
			logger.Error(err.Error())
			return err
		}
		fmt.Println(token)
		out.Result = token.AccessToken
		return nil
	}

	u := usecase.NewIOI(new(AuthInput), new(AuthOutput), handler)

	u.SetTitle("login")
	u.SetDescription("Return tokencode")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/login", u)
}

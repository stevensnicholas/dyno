package endpoints

import (
	"context"
	"dyno/internal/authentication"
<<<<<<< HEAD
=======
	"dyno/internal/logger"
>>>>>>> 13e30cc1f1abb833bc56a910660974fb688ba7e3
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
		var err error
		var code = in.Code
		var tokenAuthURL = authentication.GetTokenAuthURL(code)
		var token *authentication.Token
		if token, err = authentication.GetToken(tokenAuthURL); err != nil {
			logger.Error(err.Error())
			return err
		}
<<<<<<< HEAD

		fmt.Printf("%+v", token)
=======
>>>>>>> 13e30cc1f1abb833bc56a910660974fb688ba7e3
		out.Result = token.AccessToken
		return nil
	}

	u := usecase.NewIOI(new(AuthInput), new(AuthOutput), handler)

	u.SetTitle("login")
	u.SetDescription("Return tokencode")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/login", u)
}

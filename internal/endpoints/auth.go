package endpoints

import (
	"context"
	"time"
	
	"dyno/internal/authentication"
	"dyno/internal/logger"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type AuthInput struct {
	Code string `query:"code"`
}

type AuthOutput struct {
	Result string `json:"jwt"`
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

		var jwt string
		if jwt, err = authentication.CreateToken(time.Hour, token.AccessToken); err != nil {
			logger.Error(err.Error())
			return err
		}
		
		out.Result = jwt
		return nil
	}

	u := usecase.NewIOI(new(AuthInput), new(AuthOutput), handler)

	u.SetTitle("Authentication")
	u.SetDescription("Return token")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/login", u)
}

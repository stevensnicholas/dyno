package endpoints

import (
	"context"
	"dyno/internal/authentication"
	"dyno/internal/logger"
	"fmt"
	"time"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type AuthInput struct {
	Code string `query:"code"`
}

type AuthOutput struct {
	Result *authentication.GitHubUserInfo `json:"use info: "`
}

func Authentication(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var in = input.(*AuthInput)
		var out = output.(*AuthOutput)

		var err error
		var code = in.Code
		fmt.Println("code: ", code)
		var tokenAuthURL = authentication.GetTokenAuthURL(code)
		var token *authentication.Token
		if token, err = authentication.GetToken(tokenAuthURL); err != nil {
			logger.Error(err.Error())
			return err
		}
		fmt.Println("token: ", token)

		// var jwt string
		// if jwt, err = authentication.CreateToken(time.Hour, token.AccessToken); err != nil {
		// 	logger.Error(err.Error())
		// 	return err
		// }

		var userinfo *authentication.GitHubUserInfo
		if userinfo, err = authentication.GetUserInfo(token.AccessToken); err != nil {
			logger.Error(err.Error())
			return err
		}
		out.Result = userinfo
		return nil
	}

	u := usecase.NewIOI(new(AuthInput), new(AuthOutput), handler)

	u.SetTitle("Authentication")
	u.SetDescription("Return token")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Get("/login", u)
}

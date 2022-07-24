package endpoints

import (
	"context"
	"dyno/internal/authentication"
<<<<<<< HEAD
	"fmt"
=======
	"dyno/internal/logger"

>>>>>>> main
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
<<<<<<< HEAD

=======
>>>>>>> main
		var code = in.Code
		var tokenAuthURL = authentication.GetTokenAuthURL(code)
		var token *authentication.Token
		if token, err = authentication.GetToken(tokenAuthURL); err != nil {
<<<<<<< HEAD
			fmt.Println(err)
			return err
		}

		fmt.Printf("%+v", token)
=======
			logger.Error(err.Error())
			return err
		}

>>>>>>> main
		out.Result = token.AccessToken
		return nil
	}

	u := usecase.NewIOI(new(AuthInput), new(AuthOutput), handler)

	u.SetTitle("code")
	u.SetDescription("Return code")
	u.SetExpectedErrors(status.InvalidArgument)

	service.Post("/login", u)
}

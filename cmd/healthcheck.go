package main

import (
	"context"

	"github.com/swaggest/rest/web"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type HealthcheckInput struct{}

type HealthcheckOutput struct {
	Health string `json:"health"`
}

func healthcheck(service *web.Service) {
	handler := func(ctx context.Context, input, output interface{}) error {
		var out = output.(*HealthcheckOutput)

		out.Health = "healthy"

		return nil
	}

	u := usecase.NewIOI(new(HealthcheckInput), new(HealthcheckOutput), handler)

	u.SetTitle("Healthcheck")
	u.SetDescription("Check if the backend is healthy")
	u.SetExpectedErrors(status.Internal)

	service.Get("/healthcheck", u)
}

package main

import (
	"dyno/internal/endpoints"

	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4"
)

func registerRoutes(service *web.Service, hostdocs bool, hosthealth bool) {
	endpoints.PostEcho(service)
	endpoints.GetFuzzes(service)
	endpoints.GetFuzzesBug(service)
	endpoints.Authentication(service)
	// Swagger UI endpoint at /docs.
	if hostdocs {
		service.Docs("/docs", swgui.New)
	}

	if hosthealth {
		endpoints.Healthcheck(service)
	}
}

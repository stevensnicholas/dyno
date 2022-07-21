package main

import (
	"dyno/internal/endpoints"

	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4"
)

func registerRoutes(service *web.Service, hostdocs bool, hosthealth bool) {
	endpoints.PostEcho(service)
	endpoints.GetCode(service)
<<<<<<< HEAD

=======
	
>>>>>>> ef0003e11a62bc907e6e57d9844b69e35487828b
	// Swagger UI endpoint at /docs.
	if hostdocs {
		service.Docs("/docs", swgui.New)
	}

	if hosthealth {
		endpoints.Healthcheck(service)
	}
}

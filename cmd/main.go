package main

import (
	"flag"
	"golambda/internal/logger"
	"net/http"
	"os"
	"sort"

	"github.com/akrylysov/algnhsa"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
)

func main() {
	var gendocs = flag.Bool("gendocs", false, "Generate API Docs")
	var hostdocs = flag.Bool("hostdocs", false, "Host API Docs")
	var hosthttp = flag.Bool("http", false, "Run an HTTP server instead of in AWS Lambda")
	var healthcheck = flag.Bool("healthcheck", false, "Run healthcheck")
	var logLevel = flag.String("loglevel", "info", "Log level")

	flag.Parse()

	if *healthcheck {
		res, err := http.Get("http://localhost:8080/healthcheck")
		if err != nil || res.StatusCode != 200 {
			os.Exit(1)
		}
		os.Exit(0)
	}

	service := web.DefaultService()
	service.OpenAPI.Info.Title = "Sample Lambda API"
	service.OpenAPI.Info.WithDescription("")
	service.OpenAPI.Info.Version = "v1.0.0"

	if *gendocs {
		registerRoutes(service, *hostdocs, false)
		schemas := service.OpenAPI.Components.Schemas.MapOfSchemaOrRefValues
		for _, schema := range schemas {
			schema.Schema.Required = []string{}
			for property := range schema.Schema.Properties {
				schema.Schema.Required = append(schema.Schema.Required, property)
			}
			sort.Strings(schema.Schema.Required)
		}
		docs, err := service.OpenAPI.MarshalYAML()
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("api/openapi.yml", docs, 0600)
		if err != nil {
			panic(err)
		}
	} else if *hosthttp {
		log, err := logger.ConfigureDevelopmentLogger(*logLevel)
		if err != nil {
			panic(err)
		}
		defer log.Sync()
		service.Use(gzip.Middleware, CorsMiddleware, LoggingMiddleware)
		registerRoutes(service, *hostdocs, true)
		logger.Info("listening on port 8080")
		err = http.ListenAndServe(":8080", service)
		if err != nil {
			panic(err)
		}
	} else {
		log, err := logger.ConfigureProductionLogger(*logLevel)
		if err != nil {
			panic(err)
		}
		logger.Info("executing lambda")
		defer log.Sync()
		service.Use(APIGatewayMiddleware, CorsMiddleware, LoggingMiddleware)
		registerRoutes(service, *hostdocs, false)
		algnhsa.ListenAndServe(service, nil)
	}
}

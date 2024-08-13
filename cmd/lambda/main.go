package main

import (
	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/revotech-group/email-service/http/lambda/handlers"
	"github.com/revotech-group/go-aws/http/lambda"
	"github.com/revotech-group/go-aws/http/lambda/middleware"
	"github.com/revotech-group/go-lib/log"
)

var router *lambda.Router

func init() {
	router = lambda.NewRouter("/country", true, middleware.HttpMiddleware)

	router.Route("GET", "v1/countries/name/:name", handlers.GetByCountryNameHandler)
	router.Route("GET", "v1/countries/region/:region", handlers.GetByCountryRegionHandler)
	router.Route("GET", "v1/countries/capital/:capital", handlers.GetByCountryCapitalHandler)
	router.Route("GET", "v1/countries/lang/:lang", handlers.GetByCountryLangHandler)
	router.Route("GET", "v1/countries/currency/:currency", handlers.GetByCountryCurrencyHandler)
	router.Route("GET", "v1/countries/alpha/:alphacode", handlers.GetByCountryAlphaCodeHandler)
	router.Route("GET", "v1/countries", handlers.GetAllCountriesHandler)
}

func main() {
	log.SetupDefaultLogger(log.DefaultLevel(), true)

	awslambda.Start(router.Handler)
}

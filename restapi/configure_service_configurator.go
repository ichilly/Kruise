// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
	"github.com/ichilly/Kruise/controllers/defaults_setter/application"
	"github.com/ichilly/Kruise/controllers/renderers"
	"github.com/ichilly/Kruise/models"
	"github.com/ichilly/Kruise/restapi/operations"
	"github.com/ichilly/Kruise/restapi/operations/config"
	"github.com/ichilly/Kruise/restapi/operations/health"
)

const (
	codeRenderError = 600
)

//go:generate swagger generate server --target ../../gen --name ServiceConfigurator --spec ../../swagger.yaml

func configureFlags(api *operations.ServiceConfiguratorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ServiceConfiguratorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.TxtProducer = runtime.TextProducer()

	templateRenderer, err := renderers.NewRenderer("./templates")
	if err != nil {
		log.Fatalln(err)
	}

	api.ConfigCreateAppConfigHandler = config.CreateAppConfigHandlerFunc(
		func(params config.CreateAppConfigParams) middleware.Responder {
			if params.Application == nil {
				return config.NewCreateAppConfigBadRequest().WithPayload("application is required")
			}

			// Config app
			application.SetApplicationDefaults(params.Application)

			// Render app
			rendered, err := templateRenderer.RenderApplication(params.Application)
			if err != nil {
				errResp := &models.Error{Code: codeRenderError, Message: err.Error()}
				return config.NewCreateAppConfigDefault(500).WithPayload(errResp)
			}
			return config.NewCreateAppConfigCreated().WithPayload(rendered)
		})

	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(
		func(params health.GetHealthParams) middleware.Responder {
			return health.NewGetHealthOK().WithPayload(&models.HealthStatus{Status: "OK"})
		})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

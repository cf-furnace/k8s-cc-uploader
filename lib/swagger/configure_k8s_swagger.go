package swagger

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cf-furnace/k8s-cc-uploader/lib/swagger/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.K8sSwaggerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

// ConfigureAPI configures the CC-Uploader API server
func ConfigureAPI(api *operations.K8sSwaggerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UploadBuildArtifactsHandler = operations.UploadBuildArtifactsHandlerFunc(func(params operations.UploadBuildArtifactsParams) middleware.Responder {
		return middleware.NotImplemented("operation .UploadBuildArtifacts has not yet been implemented")
	})
	api.UploadDropletHandler = operations.UploadDropletHandlerFunc(func(params operations.UploadDropletParams) middleware.Responder {
		return middleware.NotImplemented("operation .UploadDroplet has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
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

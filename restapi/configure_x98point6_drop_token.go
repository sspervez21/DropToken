// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"98point6DropToken/droptoken"
	"98point6DropToken/restapi/operations"
)

//go:generate swagger generate server --target ..\..\98point6DropToken --name X98point6DropToken --spec ..\swagger.yml

func configureFlags(api *operations.X98point6DropTokenAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.X98point6DropTokenAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateGameHandler = operations.CreateGameHandlerFunc(func(params operations.CreateGameParams) middleware.Responder {
		return droptoken.CreateGameHandler(params)
	})
	api.GetAllGamesHandler = operations.GetAllGamesHandlerFunc(func(params operations.GetAllGamesParams) middleware.Responder {
		return droptoken.GetAllGamesHandler(params)
	})
	api.GetGameHandler = operations.GetGameHandlerFunc(func(params operations.GetGameParams) middleware.Responder {
		return droptoken.GetGameHandler(params)
	})
	api.GetAllMovesHandler = operations.GetAllMovesHandlerFunc(func(params operations.GetAllMovesParams) middleware.Responder {
		return droptoken.GetAllMovesHandler(params)
	})
	api.MakeMoveHandler = operations.MakeMoveHandlerFunc(func(params operations.MakeMoveParams) middleware.Responder {
		return droptoken.MakeMoveHandler(params)
	})
	api.GetMoveHandler = operations.GetMoveHandlerFunc(func(params operations.GetMoveParams) middleware.Responder {
		return droptoken.GetMoveHandler(params)
	})
	api.RemovePlayerHandler = operations.RemovePlayerHandlerFunc(func(params operations.RemovePlayerParams) middleware.Responder {
		return droptoken.RemovePlayerHandler(params)
	})

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

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewX98point6DropTokenAPI creates a new X98point6DropToken instance
func NewX98point6DropTokenAPI(spec *loads.Document) *X98point6DropTokenAPI {
	return &X98point6DropTokenAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		CreateGameHandler: CreateGameHandlerFunc(func(params CreateGameParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateGame has not yet been implemented")
		}),
		GetAllGamesHandler: GetAllGamesHandlerFunc(func(params GetAllGamesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAllGames has not yet been implemented")
		}),
		GetAllMovesHandler: GetAllMovesHandlerFunc(func(params GetAllMovesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAllMoves has not yet been implemented")
		}),
		GetGameHandler: GetGameHandlerFunc(func(params GetGameParams) middleware.Responder {
			return middleware.NotImplemented("operation GetGame has not yet been implemented")
		}),
		GetMoveHandler: GetMoveHandlerFunc(func(params GetMoveParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMove has not yet been implemented")
		}),
		MakeMoveHandler: MakeMoveHandlerFunc(func(params MakeMoveParams) middleware.Responder {
			return middleware.NotImplemented("operation MakeMove has not yet been implemented")
		}),
		RemovePlayerHandler: RemovePlayerHandlerFunc(func(params RemovePlayerParams) middleware.Responder {
			return middleware.NotImplemented("operation RemovePlayer has not yet been implemented")
		}),
	}
}

/*X98point6DropTokenAPI 98Point6 Drop Token microservice */
type X98point6DropTokenAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// CreateGameHandler sets the operation handler for the create game operation
	CreateGameHandler CreateGameHandler
	// GetAllGamesHandler sets the operation handler for the get all games operation
	GetAllGamesHandler GetAllGamesHandler
	// GetAllMovesHandler sets the operation handler for the get all moves operation
	GetAllMovesHandler GetAllMovesHandler
	// GetGameHandler sets the operation handler for the get game operation
	GetGameHandler GetGameHandler
	// GetMoveHandler sets the operation handler for the get move operation
	GetMoveHandler GetMoveHandler
	// MakeMoveHandler sets the operation handler for the make move operation
	MakeMoveHandler MakeMoveHandler
	// RemovePlayerHandler sets the operation handler for the remove player operation
	RemovePlayerHandler RemovePlayerHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *X98point6DropTokenAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *X98point6DropTokenAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *X98point6DropTokenAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *X98point6DropTokenAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *X98point6DropTokenAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *X98point6DropTokenAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *X98point6DropTokenAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the X98point6DropTokenAPI
func (o *X98point6DropTokenAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CreateGameHandler == nil {
		unregistered = append(unregistered, "CreateGameHandler")
	}

	if o.GetAllGamesHandler == nil {
		unregistered = append(unregistered, "GetAllGamesHandler")
	}

	if o.GetAllMovesHandler == nil {
		unregistered = append(unregistered, "GetAllMovesHandler")
	}

	if o.GetGameHandler == nil {
		unregistered = append(unregistered, "GetGameHandler")
	}

	if o.GetMoveHandler == nil {
		unregistered = append(unregistered, "GetMoveHandler")
	}

	if o.MakeMoveHandler == nil {
		unregistered = append(unregistered, "MakeMoveHandler")
	}

	if o.RemovePlayerHandler == nil {
		unregistered = append(unregistered, "RemovePlayerHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *X98point6DropTokenAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *X98point6DropTokenAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *X98point6DropTokenAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *X98point6DropTokenAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *X98point6DropTokenAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *X98point6DropTokenAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the x98point6 drop token API
func (o *X98point6DropTokenAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *X98point6DropTokenAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/drop_token"] = NewCreateGame(o.context, o.CreateGameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drop_token"] = NewGetAllGames(o.context, o.GetAllGamesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drop_token/{gameId}/moves"] = NewGetAllMoves(o.context, o.GetAllMovesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drop_token/{gameId}"] = NewGetGame(o.context, o.GetGameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drop_token/{gameId}/moves/{move_number}"] = NewGetMove(o.context, o.GetMoveHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/drop_token/{gameId}/{playerId}"] = NewMakeMove(o.context, o.MakeMoveHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/drop_token/{gameId}/{playerId}"] = NewRemovePlayer(o.context, o.RemovePlayerHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *X98point6DropTokenAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *X98point6DropTokenAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *X98point6DropTokenAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *X98point6DropTokenAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

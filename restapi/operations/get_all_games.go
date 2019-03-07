// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllGamesHandlerFunc turns a function with the right signature into a get all games handler
type GetAllGamesHandlerFunc func(GetAllGamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllGamesHandlerFunc) Handle(params GetAllGamesParams) middleware.Responder {
	return fn(params)
}

// GetAllGamesHandler interface for that can handle valid get all games params
type GetAllGamesHandler interface {
	Handle(GetAllGamesParams) middleware.Responder
}

// NewGetAllGames creates a new http.Handler for the get all games operation
func NewGetAllGames(ctx *middleware.Context, handler GetAllGamesHandler) *GetAllGames {
	return &GetAllGames{Context: ctx, Handler: handler}
}

/*GetAllGames swagger:route GET /drop_token getAllGames

GetAllGames get all games API

*/
type GetAllGames struct {
	Context *middleware.Context
	Handler GetAllGamesHandler
}

func (o *GetAllGames) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllGamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

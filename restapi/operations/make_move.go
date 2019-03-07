// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MakeMoveHandlerFunc turns a function with the right signature into a make move handler
type MakeMoveHandlerFunc func(MakeMoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MakeMoveHandlerFunc) Handle(params MakeMoveParams) middleware.Responder {
	return fn(params)
}

// MakeMoveHandler interface for that can handle valid make move params
type MakeMoveHandler interface {
	Handle(MakeMoveParams) middleware.Responder
}

// NewMakeMove creates a new http.Handler for the make move operation
func NewMakeMove(ctx *middleware.Context, handler MakeMoveHandler) *MakeMove {
	return &MakeMove{Context: ctx, Handler: handler}
}

/*MakeMove swagger:route POST /drop_token/{gameId}/{playerId} makeMove

MakeMove make move API

*/
type MakeMove struct {
	Context *middleware.Context
	Handler MakeMoveHandler
}

func (o *MakeMove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMakeMoveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

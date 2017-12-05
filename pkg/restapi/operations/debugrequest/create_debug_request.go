// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateDebugRequestHandlerFunc turns a function with the right signature into a create debug request handler
type CreateDebugRequestHandlerFunc func(CreateDebugRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDebugRequestHandlerFunc) Handle(params CreateDebugRequestParams) middleware.Responder {
	return fn(params)
}

// CreateDebugRequestHandler interface for that can handle valid create debug request params
type CreateDebugRequestHandler interface {
	Handle(CreateDebugRequestParams) middleware.Responder
}

// NewCreateDebugRequest creates a new http.Handler for the create debug request operation
func NewCreateDebugRequest(ctx *middleware.Context, handler CreateDebugRequestHandler) *CreateDebugRequest {
	return &CreateDebugRequest{Context: ctx, Handler: handler}
}

/*CreateDebugRequest swagger:route POST /debugrequest debugrequest createDebugRequest

Return a debug attachment

Return a debug attachment

*/
type CreateDebugRequest struct {
	Context *middleware.Context
	Handler CreateDebugRequestHandler
}

func (o *CreateDebugRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDebugRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
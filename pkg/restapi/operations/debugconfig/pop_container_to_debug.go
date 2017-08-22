package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PopContainerToDebugHandlerFunc turns a function with the right signature into a pop container to debug handler
type PopContainerToDebugHandlerFunc func(PopContainerToDebugParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PopContainerToDebugHandlerFunc) Handle(params PopContainerToDebugParams) middleware.Responder {
	return fn(params)
}

// PopContainerToDebugHandler interface for that can handle valid pop container to debug params
type PopContainerToDebugHandler interface {
	Handle(PopContainerToDebugParams) middleware.Responder
}

// NewPopContainerToDebug creates a new http.Handler for the pop container to debug operation
func NewPopContainerToDebug(ctx *middleware.Context, handler PopContainerToDebugHandler) *PopContainerToDebug {
	return &PopContainerToDebug{Context: ctx, Handler: handler}
}

/*PopContainerToDebug swagger:route DELETE /debugconfig/platform/containers/{node}/latest debugconfig popContainerToDebug

Pop the latest debugconfig for a node in the cluster - this is used by the squash client.

*/
type PopContainerToDebug struct {
	Context *middleware.Context
	Handler PopContainerToDebugHandler
}

func (o *PopContainerToDebug) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPopContainerToDebugParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

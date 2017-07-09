package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPatientHandlerFunc turns a function with the right signature into a get patient handler
type GetPatientHandlerFunc func(GetPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPatientHandlerFunc) Handle(params GetPatientParams) middleware.Responder {
	return fn(params)
}

// GetPatientHandler interface for that can handle valid get patient params
type GetPatientHandler interface {
	Handle(GetPatientParams) middleware.Responder
}

// NewGetPatient creates a new http.Handler for the get patient operation
func NewGetPatient(ctx *middleware.Context, handler GetPatientHandler) *GetPatient {
	return &GetPatient{Context: ctx, Handler: handler}
}

/*GetPatient swagger:route GET /patient getPatient

GetPatient get patient API

*/
type GetPatient struct {
	Context *middleware.Context
	Handler GetPatientHandler
}

func (o *GetPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

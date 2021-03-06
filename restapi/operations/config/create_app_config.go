// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateAppConfigHandlerFunc turns a function with the right signature into a create app config handler
type CreateAppConfigHandlerFunc func(CreateAppConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAppConfigHandlerFunc) Handle(params CreateAppConfigParams) middleware.Responder {
	return fn(params)
}

// CreateAppConfigHandler interface for that can handle valid create app config params
type CreateAppConfigHandler interface {
	Handle(CreateAppConfigParams) middleware.Responder
}

// NewCreateAppConfig creates a new http.Handler for the create app config operation
func NewCreateAppConfig(ctx *middleware.Context, handler CreateAppConfigHandler) *CreateAppConfig {
	return &CreateAppConfig{Context: ctx, Handler: handler}
}

/*CreateAppConfig swagger:route POST /v1/app/configs config createAppConfig

Creates the application configurations

*/
type CreateAppConfig struct {
	Context *middleware.Context
	Handler CreateAppConfigHandler
}

func (o *CreateAppConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAppConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

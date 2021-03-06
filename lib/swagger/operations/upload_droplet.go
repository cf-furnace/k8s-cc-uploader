package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UploadDropletHandlerFunc turns a function with the right signature into a upload droplet handler
type UploadDropletHandlerFunc func(UploadDropletParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadDropletHandlerFunc) Handle(params UploadDropletParams) middleware.Responder {
	return fn(params)
}

// UploadDropletHandler interface for that can handle valid upload droplet params
type UploadDropletHandler interface {
	Handle(UploadDropletParams) middleware.Responder
}

// NewUploadDroplet creates a new http.Handler for the upload droplet operation
func NewUploadDroplet(ctx *middleware.Context, handler UploadDropletHandler) *UploadDroplet {
	return &UploadDroplet{Context: ctx, Handler: handler}
}

/*UploadDroplet swagger:route POST /droplet/{guid} uploadDroplet

UploadDroplet upload droplet API

*/
type UploadDroplet struct {
	Context *middleware.Context
	Handler UploadDropletHandler
}

func (o *UploadDroplet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUploadDropletParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

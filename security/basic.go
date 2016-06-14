package main

import (
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}

			// Proceed
			goa.LogInfo(ctx, "auth", "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

// BasicController implements the BasicAuth resource.
type BasicController struct {
	*goa.Controller
}

// NewBasicController creates a BasicAuth controller.
func NewBasicController(service *goa.Service) *BasicController {
	return &BasicController{Controller: service.NewController("BasicController")}
}

// Secured runs the secured action.
func (c *BasicController) Secured(ctx *app.SecuredBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecured runs the unsecured action.
func (c *BasicController) Unsecured(ctx *app.UnsecuredBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}
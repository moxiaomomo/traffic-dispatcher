package client

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	admin "path/to/service/proto/admin"
)

type adminKey struct {}

// FromContext retrieves the client from the Context
func AdminFromContext(ctx context.Context) (admin.AdminService, bool) {
	c, ok := ctx.Value(adminKey{}).(admin.AdminService)
	return c, ok
}

// Client returns a wrapper for the AdminClient
func AdminWrapper(service micro.Service) server.HandlerWrapper {
	client := admin.NewAdminService("go.micro.service.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, adminKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}

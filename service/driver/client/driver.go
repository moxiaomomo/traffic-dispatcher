package client

import (
	"context"

	"traffic-dispatcher/proto/lbs"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type driverKey struct{}

// FromContext retrieves the client from the Context
func DriverFromContext(ctx context.Context) (lbs.LbsService, bool) {
	c, ok := ctx.Value(driverKey{}).(lbs.LbsService)
	return c, ok
}

// Client returns a wrapper for the DriverClient
func DriverWrapper(service micro.Service) server.HandlerWrapper {
	client := lbs.NewLbsService("go.micro.service.lbs", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, driverKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}

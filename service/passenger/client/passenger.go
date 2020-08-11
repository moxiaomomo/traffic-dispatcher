package client

import (
	"context"

	"traffic-dispatcher/proto/lbs"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type passengerKey struct{}

// FromContext retrieves the client from the Context
func PassengerFromContext(ctx context.Context) (lbs.LbsService, bool) {
	c, ok := ctx.Value(passengerKey{}).(lbs.LbsService)
	return c, ok
}

// Client returns a wrapper for the PassengerClient
func PassengerWrapper(service micro.Service) server.HandlerWrapper {
	client := lbs.NewLbsService("go.micro.service.lbs", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, passengerKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}

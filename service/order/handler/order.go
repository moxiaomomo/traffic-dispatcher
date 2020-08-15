package handler

import (
	"context"
	hello "traffic-dispatcher/proto/hello"

	"github.com/micro/go-micro/v2/logger"
)

type Order struct{}

func (o *Order) Hello(ctx context.Context, req *hello.SayRequest, rsp *hello.SayResponse) error {
	logger.Info("Received Hello request:", req.GetFrom())

	rsp.Msg = "Echo hello from order backend service"
	return nil
}

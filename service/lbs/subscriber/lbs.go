package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	"traffic-dispatcher/proto/lbs"
)

type Lbs struct{}

func (e *Lbs) Handle(ctx context.Context, msg *lbs.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *lbs.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}

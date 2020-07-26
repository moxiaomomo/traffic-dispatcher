package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	api "user/api/proto/api"
)

type Api struct{}

func (e *Api) Handle(ctx context.Context, msg *api.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *api.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}

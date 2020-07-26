package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	api "user/api/proto/api"
)

type Api struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Api) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Api.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Api) Stream(ctx context.Context, req *api.StreamingRequest, stream api.Api_StreamStream) error {
	log.Infof("Received Api.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&api.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Api) PingPong(ctx context.Context, stream api.Api_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&api.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

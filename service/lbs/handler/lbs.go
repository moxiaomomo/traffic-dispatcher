package handler

import (
	"context"
	"traffic-dispatcher/proto/lbs"

	log "github.com/micro/go-micro/v2/logger"
)

type Lbs struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Lbs) Call(ctx context.Context, req *lbs.Request, rsp *lbs.Response) error {
	log.Info("Received Lbs.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Lbs) Stream(ctx context.Context, req *lbs.StreamingRequest, stream lbs.Lbs_StreamStream) error {
	log.Infof("Received Lbs.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&lbs.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Lbs) PingPong(ctx context.Context, stream lbs.Lbs_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&lbs.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

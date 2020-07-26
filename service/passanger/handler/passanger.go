package handler

import (
	"context"
	"encoding/json"
	"traffic-dispatcher/proto/lbs"

	log "github.com/micro/go-micro/v2/logger"

	"passanger/client"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

type Passanger struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Passanger.Call is called by the API as /passanger/call with post body {"name": "foo"}
func (e *Passanger) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Passanger.Call request")

	// extract the client from the context
	passangerClient, ok := client.PassangerFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.passanger.passanger.call", "passanger client not found")
	}

	// make request
	response, err := passangerClient.Call(ctx, &lbs.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.passanger.passanger.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

// Passanger.GetLocation is called by the API as /passanger/getLocation with post body {"code": "0"}
func (e *Passanger) GetLocation(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Passanger.GetLocation request")
	rsp.StatusCode = 200
	rsp.Body = `{"code": "0"}`

	return nil
}

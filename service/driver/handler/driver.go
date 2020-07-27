package handler

import (
	"context"
	"encoding/json"

	log "github.com/micro/go-micro/v2/logger"

	"driver/client"
	"traffic-dispatcher/proto/lbs"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

type Driver struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Driver.Call is called by the API as /driver/call with post body {"name": "foo"}
func (e *Driver) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Driver.Call request")

	// extract the client from the context
	driverClient, ok := client.DriverFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.driver.driver.call", "driver client not found")
	}

	// make request
	response, err := driverClient.Call(ctx, &lbs.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.driver.driver.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

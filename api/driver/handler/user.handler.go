package handler

import (
	"context"
	"encoding/json"

	hello "traffic-dispatcher/proto/hello"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/logger"
)

type Say struct {
	Client hello.SayService
}

//http://localhost:8080/driver/say/hello?name=xxxx
func (s *Say) Hello(ctx context.Context, req map[string]string, rsp *api.Response) error {
//func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	logger.Info("Received Say.Hello API request")

	name, ok := req["name"]
	if !ok || len(name) == 0 {
		return errors.BadRequest("go.micro.api.driver", "Name cannot be blank")
	}

	// 在restful api中调用rpc服务
	response, err := s.Client.Hello(ctx, &hello.SayRequest{
		From: "hahah",
		To:   "xxx",
		Msg:  "Request from client",
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"from": response.GetFrom(),
		"to":   response.GetTo(),
		"Msg":  response.GetMsg(),
	})
	rsp.Body = string(b)

	return nil
}

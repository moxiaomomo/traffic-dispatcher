package handler

import (
	"context"
	"encoding/json"
	"time"
	"traffic-dispatcher/model"
	order "traffic-dispatcher/proto/order"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

type Order struct {
	Client order.OrderService
}

func parseReqOrderBody(req *api.Request) (reqOrder order.Order, err error) {
	err = json.Unmarshal([]byte(req.Body), &reqOrder)
	return
}

func (s *Order) CreateOrder(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqOrder, err := parseReqOrderBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
	}
	reqOrder.CreateAt = time.Now().Unix()
	reqOrder.Status = int32(model.OrderCreated)

	response, err := s.Client.CreateOrder(ctx, &order.ReqCreateOrder{
		Order: &reqOrder,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.GetCode(),
		"order": response.GetOrder(),
		"msg":   response.GetMessage(),
	})
	rsp.Body = string(b)

	return nil
}

func (s *Order) ConfirmGetOn(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqOrder, err := parseReqOrderBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
	}
	reqOrder.CreateAt = time.Now().Unix()
	reqOrder.Status = int32(model.OrderCreated)

	response, err := s.Client.ConfirmGetOn(ctx, &order.ReqConfirmGetOn{
		Order: &reqOrder,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.GetCode(),
		"order": response.GetOrder(),
		"msg":   response.GetMessage(),
	})
	rsp.Body = string(b)

	return nil
}

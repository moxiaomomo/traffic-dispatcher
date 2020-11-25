package handler

import (
	"context"
	"encoding/json"
	"time"
	"traffic-dispatcher/config"
	"traffic-dispatcher/model"
	order "traffic-dispatcher/proto/order"

	"traffic-dispatcher/api/passenger/mq"

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

	if response.GetCode() == int32(config.StatusOK) {
		// 发布订单消息
		respOrder := response.GetOrder()
		respOrder.SrcAddr = reqOrder.SrcAddr
		respOrder.DestAddr = reqOrder.DestAddr
		orderJSON, _ := json.Marshal(respOrder)
		mq.Publish(config.DriverLbsMQTopic, string(orderJSON))
	}

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

// CancelOrder todo in development
func (s *Order) CancelOrder(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqOrder, err := parseReqOrderBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
	}
	reqOrder.CreateAt = time.Now().Unix()
	reqOrder.Status = int32(model.OrderCreated)

	response, err := s.Client.CancelOrder(ctx, &order.ReqCancelOrder{
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

// QueryOrderHis todo in development
func (s *Order) QueryOrderHis(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var hisReq order.ReqOrderHis
	err := json.Unmarshal([]byte(req.Body), &hisReq)
	if err != nil || len(hisReq.UserId) == 0 {
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
	}

	response, err := s.Client.QueryOrderHis(ctx, &hisReq)
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":   response.GetCode(),
		"orders": response.GetOrders(),
		"msg":    response.GetMessage(),
	})
	rsp.Body = string(b)

	// TODO ceshi
	// 发布订单消息
	mq.Publish(config.DriverLbsMQTopic, string(b))
	return nil
}

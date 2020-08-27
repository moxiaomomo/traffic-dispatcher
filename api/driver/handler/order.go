package handler

import (
	"context"
	"encoding/json"
	"time"
	"traffic-dispatcher/model"
	"traffic-dispatcher/proto/order"

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

func (s *Order) AcceptOrder(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqOrder, err := parseReqOrderBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.driver", "request invalid")
	}
	reqOrder.AcceptAt = time.Now().Unix()
	reqOrder.Status = int32(model.OrderAccepted)

	response, err := s.Client.AcceptOrder(ctx, &order.ReqAcceptOrder{
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

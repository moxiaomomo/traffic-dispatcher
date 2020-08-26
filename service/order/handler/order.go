package handler

import (
	"context"
	"traffic-dispatcher/config"
	dbmysql "traffic-dispatcher/db"
	order "traffic-dispatcher/proto/order"
	"traffic-dispatcher/util"

	"github.com/micro/go-micro/v2/logger"
)

type Order struct{}

func (o *Order) CreateOrder(ctx context.Context, req *order.ReqCreateOrder, rsp *order.RespCreateOrder) error {
	logger.Infof("Received CreateOrder request: %s\n", req.GetOrder().GetPassengerId())

	if dbOrder, err := dbmysql.CreateOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

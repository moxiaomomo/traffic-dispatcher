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
	logger.Infof("Received CreateOrder request: %s\n", req.GetOrder().PassengerId)

	if dbOrder, err := dbmysql.CreateOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

func (o *Order) AcceptOrder(ctx context.Context, req *order.ReqAcceptOrder, rsp *order.RespAcceptOrder) error {
	logger.Infof("Received AcceptOrder request: %s\n", req.GetOrder().DriverId)

	if dbOrder, err := dbmysql.AcceptOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

func (o *Order) ConfirmGetOn(ctx context.Context, req *order.ReqConfirmGetOn, rsp *order.RespConfirmGetOn) error {
	logger.Infof("Received ConfirmGetOn request: %s\n", req.GetOrder().DriverId)

	if dbOrder, err := dbmysql.ConfirmGetOn(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

func (o *Order) StartOrder(ctx context.Context, req *order.ReqStartOrder, rsp *order.RespStartOrder) error {
	logger.Infof("Received StartOrder request: %s\n", req.GetOrder().DriverId)

	if dbOrder, err := dbmysql.StartOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

func (o *Order) FinishOrder(ctx context.Context, req *order.ReqFinishOrder, rsp *order.RespFinishOrder) error {
	logger.Infof("Received FinishOrder request: %s\n", req.GetOrder().DriverId)

	if dbOrder, err := dbmysql.FinishOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
		rsp.Code = int32(config.StatusOK)
		rsp.Order = util.OrmOrder2ProtoOrder(dbOrder)
	} else {
		rsp.Code = int32(config.StatusServerError)
	}
	return nil
}

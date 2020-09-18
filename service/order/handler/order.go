package handler

import (
	"context"
	"traffic-dispatcher/config"
	dbmysql "traffic-dispatcher/db"
	"traffic-dispatcher/model"
	"traffic-dispatcher/model/orm"
	order "traffic-dispatcher/proto/order"
	"traffic-dispatcher/util"

	"github.com/micro/go-micro/v2/logger"
)

type Order struct{}

func (o *Order) CreateOrder(ctx context.Context, req *order.ReqCreateOrder, rsp *order.RespCreateOrder) error {
	logger.Infof("Received CreateOrder request: %s\n", req.GetOrder().PassengerId)

	if _, exist := dbmysql.QueryActiveOrder(req.GetOrder().PassengerId); exist {
		rsp.Code = int32(config.StatusDulplicated)
		return nil
	}

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

func (o *Order) CancelOrder(ctx context.Context, req *order.ReqCancelOrder, rsp *order.RespCancelOrder) error {
	logger.Infof("Received CancelOrder request: %s\n", req.GetOrder().DriverId)

	if nowOrder, err := dbmysql.QueryOrder(req.GetOrder().GetOrderId()); err != nil {
		rsp.Code = int32(config.StatusParamInvalid)
		return nil
	} else if nowOrder.Status != int32(model.OrderCreated) && nowOrder.Status != int32(model.OrderAccepted) {
		rsp.Code = int32(config.StatusNotPermitted)
		rsp.Message = "当前订单不满足取消条件"
		return nil
	}

	if dbOrder, err := dbmysql.CancelOrder(util.ProtoOrder2OrmOrder(req.GetOrder())); err == nil {
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

// QueryOrderHis 查询行程订单历史
func (o *Order) QueryOrderHis(ctx context.Context, req *order.ReqOrderHis, rsp *order.RespOrderHis) error {
	logger.Infof("Received QueryOrderHis request: %s\n", req.GetUserId())

	var dbOrders []orm.Order
	var err error
	if req.GetRole() == int32(model.ClientDriver) {
		dbOrders, err = dbmysql.QueryOrderByDriver(req.UserId, req.FromTS, req.ToTS)
	} else if req.GetRole() == int32(model.ClientPassenger) {
		dbOrders, err = dbmysql.QueryOrderByPassenger(req.UserId, req.FromTS, req.ToTS)
	}

	if err != nil {
		rsp.Code = int32(config.StatusServerError)
		return nil
	}

	var ptOrders []*order.Order
	for _, dbOrder := range dbOrders {
		ptOrders = append(ptOrders, util.OrmOrder2ProtoOrder(&dbOrder))
	}

	rsp.Orders = ptOrders
	return nil
}

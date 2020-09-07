package db

import (
	"errors"
	"fmt"
	"time"

	dbmysql "traffic-dispatcher/db/mysql"
	"traffic-dispatcher/model/orm"
	"traffic-dispatcher/util"
)

func genOrderID(order *orm.Order) string {
	if len(order.OrderId) == 32 {
		return order.OrderId
	}

	ts := fmt.Sprintf("%x", time.Now().Unix())
	tmp := fmt.Sprintf("%s%s%s", order.PassengerId, order.DriverId, ts)
	return util.MD5([]byte(tmp))
}

// CreateOrder 创建订单
func CreateOrder(order *orm.Order) (*orm.Order, error) {
	order.OrderId = genOrderID(order)
	// order.Status = int32(model.OrderCreated)
	err := dbmysql.Conn().Create(order).Error
	return order, err
}

// AcceptOrder 接收订单
func AcceptOrder(order *orm.Order) (*orm.Order, error) {
	if order.OrderId == "" || order.DriverId == "" {
		return nil, errors.New("Invalid order to accept")
	}

	// 更新指定字段
	err := dbmysql.Conn().Model(&order).Where("order_id = ?", order.OrderId).Updates(
		orm.Order{
			AcceptAt: order.AcceptAt,
			DriverId: order.DriverId,
			Status:   order.Status,
		},
	).Error
	return order, err
}

// ConfirmGetOn 处理订单(确认上车)
func ConfirmGetOn(order *orm.Order) (*orm.Order, error) {
	if order.OrderId == "" {
		return nil, errors.New("Invalid order to confirm geton")
	}

	// 更新指定字段
	err := dbmysql.Conn().Model(&order).Where("order_id = ?", order.OrderId).Updates(
		orm.Order{
			GetOnAt: order.GetOnAt,
			Status:  order.Status,
		},
	).Error
	return order, err
}

// StartOrder 处理订单(开始行程)
func StartOrder(order *orm.Order) (*orm.Order, error) {
	if order.OrderId == "" {
		return nil, errors.New("Invalid order to start")
	}

	// 更新指定字段
	err := dbmysql.Conn().Model(&order).Where("order_id = ?", order.OrderId).Updates(
		orm.Order{
			StartAt: order.StartAt,
			Status:  order.Status,
		},
	).Error
	return order, err
}

// FinishOrder 处理订单(开始行程)
func FinishOrder(order *orm.Order) (*orm.Order, error) {
	if order.OrderId == "" {
		return nil, errors.New("Invalid order to finish")
	}

	// 更新指定字段
	err := dbmysql.Conn().Model(&order).Where("order_id = ?", order.OrderId).Updates(
		orm.Order{
			FinishAt: order.FinishAt,
			Status:   order.Status,
		},
	).Error
	return order, err
}

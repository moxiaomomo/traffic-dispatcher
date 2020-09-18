package db

import (
	"errors"
	"fmt"
	"log"
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

// QueryOrder 根据订单id查询订单
func QueryOrder(orderID string) (order orm.Order, err error) {
	dbmysql.Conn().Model(&order).Where("order_id = ?", orderID).First(&order)
	if order.OrderId != orderID {
		err = errors.New("No order matched")
		return
	}
	return
}

// QueryActiveOrder 根据乘客id查询未完成订单
func QueryActiveOrder(pID string) (order orm.Order, exist bool) {
	dbmysql.Conn().Model(&order).Where("passenger_id = ? and status not in (3,4)", pID).First(&order)
	if order.PassengerId != pID {
		exist = false
		return
	}
	exist = true
	return
}

// QueryOrderByDriver 根据司机id及日期范围来查询订单
func QueryOrderByDriver(userID string, fromTS int64, toTS int64) (orders []orm.Order, err error) {
	if toTS > fromTS {
		from := time.Unix(fromTS, 0)
		to := time.Unix(toTS, 0)
		dbmysql.Conn().Where("driver_id = ? and create_at>=? and create_at<=?", userID, from, to).Limit(30).Find(&orders)
	} else {
		dbmysql.Conn().Where("driver_id = ?", userID).Limit(30).Find(&orders)
	}

	return
}

// QueryOrderByPassenger 根据乘客id及日期范围来查询订单
func QueryOrderByPassenger(userID string, fromTS int64, toTS int64) (orders []orm.Order, err error) {
	if toTS > fromTS {
		from := time.Unix(fromTS, 0)
		to := time.Unix(toTS, 0)
		dbmysql.Conn().Where("passenger_id = ? and create_at>=? and create_at<=?", userID, from, to).Limit(30).Find(&orders)
	} else {
		dbmysql.Conn().Where("passenger_id = ?", userID).Limit(30).Find(&orders)
	}
	log.Print(userID)
	return
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
			GetonAt: order.GetonAt,
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

// CancelOrder 取消订单
func CancelOrder(order *orm.Order) (*orm.Order, error) {
	if order.OrderId == "" {
		return nil, errors.New("Invalid order to cancel")
	}

	// 更新指定字段
	err := dbmysql.Conn().Model(&order).Where("order_id = ?", order.OrderId).Updates(
		orm.Order{
			CancelAt: order.CancelAt,
			Status:   order.Status,
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

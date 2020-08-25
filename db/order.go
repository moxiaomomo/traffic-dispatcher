package db

import (
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
func CreateOrder(order *orm.Order) error {
	order.OrderId = genOrderID(order)
	if err := dbmysql.Conn().Create(order).Error; err != nil {
		return err
	}
	return nil
}

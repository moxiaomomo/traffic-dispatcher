package orm

import (
	"time"
)

// Order 对应table: tbl_order
type Order struct {
	// Id 用户id
	Id uint
	// OrderId 订单id
	OrderId string
	// SrcGeo 起始位置
	SrcGeo string
	// DestGeo 目的位置
	DestGeo string
	// CreateAt 创建时间
	CreateAt *time.Time
	// Accept 接单时间
	AcceptAt *time.Time
	// Cancel 取消时间
	CancelAt *time.Time
	// FinishAt 完成时间
	FinishAt *time.Time
	// CancelRole 发起取消操作的角色
	CancelRole int32
	// Cost 订单交易价格
	Cost float64
	// PassengerId 乘客Id
	PassengerId string
	// DriverId 司机Id
	DriverId string
	// Status 订单状态
	Status int32
}

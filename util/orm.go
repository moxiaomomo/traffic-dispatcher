package util

import (
	"time"
	"traffic-dispatcher/model/orm"
	orderProto "traffic-dispatcher/proto/order"
	userProto "traffic-dispatcher/proto/user"
)

// OrmUser2ProtoUser *orm.User->userProto.User
func OrmUser2ProtoUser(user *orm.User) *userProto.User {
	tmp := userProto.User{
		Id:             int64(user.Id),
		Role:           int32(user.Role),
		UserID:         user.UserId,
		UserName:       user.UserName,
		UserPwd:        user.UserPwd,
		Email:          user.Email,
		Phone:          user.Phone,
		EmailValidated: user.EmailValidated,
		PhoneValidated: user.PhoneValidated,
		Profile:        user.Profile,
		Status:         int32(user.Status),
		// Token:          user.Token,
	}
	if user.SignupAt != nil {
		tmp.SignupAt = uint64(user.SignupAt.Unix())
	}
	if user.LastActive != nil {
		tmp.LastActive = uint64(user.LastActive.Unix())
	}
	return &tmp
}

// ProtoUser2OrmUser *userProto.User->orm.User
func ProtoUser2OrmUser(user *userProto.User) *orm.User {
	tmp := orm.User{
		Id:             uint(user.Id),
		Role:           int(user.Role),
		UserId:         user.UserID,
		UserName:       user.UserName,
		UserPwd:        user.UserPwd,
		Email:          user.Email,
		Phone:          user.Phone,
		EmailValidated: user.EmailValidated,
		PhoneValidated: user.PhoneValidated,
		Profile:        user.Profile,
		Status:         int(user.Status),
		// Token:          user.Token,
	}
	if user.SignupAt > 0 {
		signupAt := time.Unix(int64(user.SignupAt), 0)
		tmp.SignupAt = &signupAt
	}
	if user.LastActive > 0 {
		lastActive := time.Unix(int64(user.LastActive), 0)
		tmp.LastActive = &lastActive
	}
	return &tmp
}

// OrmOrder2ProtoOrder orm order to proto order
func OrmOrder2ProtoOrder(order *orm.Order) *orderProto.Order {
	tmp := orderProto.Order{
		Id:          int64(order.Id),
		OrderId:     order.OrderId,
		SrcGeo:      order.SrcGeo,
		DestGeo:     order.DestGeo,
		CancelRole:  order.CancelRole,
		Cost:        order.Cost,
		PassengerId: order.PassengerId,
		DriverId:    order.DriverId,
		Status:      order.Status,
	}
	if order.CreateAt != nil {
		tmp.CreateAt = order.CreateAt.Unix()
	}
	if order.AcceptAt != nil {
		tmp.AcceptAt = order.AcceptAt.Unix()
	}
	if order.GetOnAt != nil {
		tmp.GetOnAt = order.GetOnAt.Unix()
	}
	if order.StartAt != nil {
		tmp.StartAt = order.StartAt.Unix()
	}
	if order.CancelAt != nil {
		tmp.CancelAt = order.CancelAt.Unix()
	}
	if order.FinishAt != nil {
		tmp.FinishAt = order.FinishAt.Unix()
	}
	return &tmp
}

// ProtoOrder2OrmOrder proto order to orm proto
func ProtoOrder2OrmOrder(order *orderProto.Order) *orm.Order {
	tmp := orm.Order{
		Id:          uint(order.Id),
		OrderId:     order.OrderId,
		SrcGeo:      order.SrcGeo,
		DestGeo:     order.DestGeo,
		CancelRole:  order.CancelRole,
		Cost:        order.Cost,
		PassengerId: order.PassengerId,
		DriverId:    order.DriverId,
		Status:      order.Status,
	}
	if order.CreateAt > 0 {
		createAt := time.Unix(order.CreateAt, 0)
		tmp.CreateAt = &createAt
	}
	if order.AcceptAt > 0 {
		acceptAt := time.Unix(order.AcceptAt, 0)
		tmp.AcceptAt = &acceptAt
	}
	if order.GetOnAt > 0 {
		getOnAt := time.Unix(order.GetOnAt, 0)
		tmp.GetOnAt = &getOnAt
	}
	if order.StartAt > 0 {
		startAt := time.Unix(order.StartAt, 0)
		tmp.StartAt = &startAt
	}
	if order.CancelAt > 0 {
		cancelAt := time.Unix(order.CancelAt, 0)
		tmp.CancelAt = &cancelAt
	}
	if order.FinishAt > 0 {
		finishAt := time.Unix(order.FinishAt, 0)
		tmp.FinishAt = &finishAt
	}
	return &tmp
}

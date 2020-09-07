package util

import (
	"fmt"
	"reflect"
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

	convertTimeField(&tmp, user)

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

	convertTimeField(&tmp, user)

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

	convertTimeField(&tmp, order)

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

	convertTimeField(&tmp, order)

	return &tmp
}

func struct2map(ptr interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	names := []string{"SignupAt", "LastActive", "CreateAt", "AcceptAt",
		"GetOnAt", "StartAt", "CancelAt", "FinishAt"}

	tv := reflect.ValueOf(ptr).Elem()

	for i := 0; i < tv.NumField(); i++ {
		fieldInfo := tv.Type().Field(i)
		tt := fmt.Sprintf("%v", fieldInfo.Type)
		if yes, _ := Contain(names, fieldInfo.Name); yes {
			val := tv.FieldByName(fieldInfo.Name).Interface()
			if tt == "*time.Time" {
				if ok := val.(*time.Time); ok != nil {
					res[fieldInfo.Name] = ok.Unix()
				}
			} else {
				if ok := val.(int64); ok > 0 {
					nTime := time.Unix(ok, 0)
					res[fieldInfo.Name] = &nTime
				}
			}
		}
	}
	return res
}

func convertTimeField(ptrDest interface{}, ptrSrc interface{}) {
	fields := struct2map(ptrSrc)
	v := reflect.ValueOf(ptrDest).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		if value, ok := fields[fieldInfo.Name]; ok {
			//给结构体赋值 保证赋值时数据类型一致
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
}

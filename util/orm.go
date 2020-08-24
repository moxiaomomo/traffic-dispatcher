package util

import (
	"time"
	"traffic-dispatcher/model/orm"
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
		Token:          user.Token,
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
		Token:          user.Token,
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

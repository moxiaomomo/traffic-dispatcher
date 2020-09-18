package handler

import (
	"context"
	dbproxy "traffic-dispatcher/db"
	"traffic-dispatcher/model/orm"

	"github.com/micro/go-micro/v2/logger"

	user "traffic-dispatcher/proto/user"
	"traffic-dispatcher/util"
)

type User struct{}

// Signup 用户注册
func (e *User) Signup(ctx context.Context, req *user.ReqSignup, rsp *user.RespSignup) error {
	logger.Infof("user signup: %s\n", req.User.GetUserName())

	dbUser := orm.User{
		Role:     int(req.User.Role),
		UserName: req.User.UserName,
		UserPwd:  req.User.UserPwd,
		Status:   0,
	}
	err := dbproxy.Signup(&dbUser)
	if err == nil {
		rsp.Code = 1
		rsp.Message = "Signup succeeded."
	} else {
		rsp.Code = -1
		rsp.Message = "Signup failed: " + err.Error()
	}

	return nil
}

// Signin 用户登录
func (e *User) Signin(ctx context.Context, req *user.ReqSignin, rsp *user.RespSignin) error {
	logger.Infof("user signin: %s\n", req.User.GetUserName())

	dbUser := util.ProtoUser2OrmUser(req.User)
	rspUser, token, err := dbproxy.Signin(dbUser)
	if err == nil {
		rsp.Code = 1
		pUser := util.OrmUser2ProtoUser(&rspUser)
		pUser.Token = token
		rsp.User = pUser
		rsp.Message = "Signin succeeded."
	} else {
		rsp.Code = -1
		rsp.Message = "Signin failed: " + err.Error()
	}
	return nil
}

// Signout 用户注销登录
func (e *User) Signout(ctx context.Context, req *user.ReqSignout, rsp *user.RespSignout) error {
	// TODO do something
	return nil
}

// UserInfo 用户注册
func (e *User) UserInfo(ctx context.Context, req *user.ReqUserInfo, rsp *user.RespUserInfo) error {
	logger.Infof("user info query: %s\n", req.User.GetUserName())
	return nil
}

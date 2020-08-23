package handler

import (
	"context"
	"strconv"
	dbproxy "traffic-dispatcher/db"
	"traffic-dispatcher/model/orm"

	"github.com/micro/go-micro/v2/logger"

	user "traffic-dispatcher/proto/user"
)

type User struct{}

// QueryUserByName 实现了user.pb.micro.go中的UserHandler接口
func (e *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	//rsp.User.Name = "Hello " + req.UserName//rsp.User是零值（nil），不能直接对其属性赋值，所以需要创建新对象赋值
	ID64, _ := strconv.ParseInt(req.UserID, 10, 64)
	rsp.User = &user.User{
		Id:   ID64,
		Name: req.UserName,
		Pwd:  req.UserPwd,
	}
	rsp.Success = true
	return nil
}

// Signup 用户注册
func (e *User) Signup(ctx context.Context, req *user.ReqSignup, rsp *user.RespSignup) error {
	logger.Infof("user signup: %s\n", req.GetUsername())

	dbUser := orm.User{
		Role:     int(req.GetRole()),
		UserName: req.GetUsername(),
		UserPwd:  req.GetPassword(),
		Status:   0,
	}
	err := dbproxy.Signup(&dbUser)
	if err == nil {
		rsp.Code = 1
		rsp.Message = "Signup succeeded."
	} else {
		rsp.Code = 0
		rsp.Message = "Signup failed: " + err.Error()
	}

	return nil
}

// Signin 用户登录
func (e *User) Signin(ctx context.Context, req *user.ReqSignin, rsp *user.RespSignin) error {
	logger.Infof("user signin: %s\n", req.GetUsername())

	dbUser := orm.User{
		Role:     int(req.GetRole()),
		UserName: req.GetUsername(),
		UserPwd:  req.GetPassword(),
	}
	err := dbproxy.Signin(&dbUser)
	if err == nil {
		rsp.Code = 1
		rsp.Message = "Signin succeeded."
	} else {
		rsp.Code = 0
		rsp.Message = "Signin failed: " + err.Error()
	}
	return nil
}

// UserInfo 用户注册
func (e *User) UserInfo(ctx context.Context, req *user.ReqUserInfo, rsp *user.RespUserInfo) error {
	logger.Infof("user info query: %s\n", req.GetUsername())
	return nil
}

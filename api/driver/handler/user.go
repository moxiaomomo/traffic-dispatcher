package handler

import (
	"context"
	"encoding/json"

	user "traffic-dispatcher/proto/user"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

type User struct {
	Client user.UserService
}

func parseReqBody(req *api.Request) (reqUser user.User, err error) {
	err = json.Unmarshal([]byte(req.Body), &reqUser)
	return
}

// func (s *Say) Hello(ctx context.Context, req map[string]string, rsp *api.Response) error {
func (s *User) QueryUserByName(ctx context.Context, req *api.Request, rsp *api.Response) error {
	// logger.Info("Received User.QueryUserByName API request")
	// name, ok := req["name"]
	// if !ok || len(name) == 0 {
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.driver", "Name cannot be blank")
	}

	var reqUser = user.User{
		UserName: name.Values[0],
	}
	// 在restful api中调用rpc服务
	response, err := s.Client.UserInfo(ctx, &user.ReqUserInfo{
		User: &reqUser,
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.GetCode(),
		"user": response.GetUser(),
		"msg":  response.GetMessage(),
	})
	rsp.Body = string(b)

	return nil
}

func (s *User) Signup(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqUser, err := parseReqBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.driver", "request invalid")
	}

	response, err := s.Client.Signup(ctx, &user.ReqSignup{
		User: &reqUser,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.GetCode(),
		"msg":  response.GetMessage(),
	})
	rsp.Body = string(b)

	return nil
}

func (s *User) Signin(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqUser, err := parseReqBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.driver", "request invalid")
	}

	response, err := s.Client.Signin(ctx, &user.ReqSignin{
		User: &reqUser,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.GetCode(),
		"user": response.GetUser(),
		"msg":  response.GetMessage(),
	})
	rsp.Body = string(b)

	return nil
}

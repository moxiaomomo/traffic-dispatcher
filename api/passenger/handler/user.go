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

func (s *User) Signup(ctx context.Context, req *api.Request, rsp *api.Response) error {
	username, ok := req.Get["username"]
	if !ok || len(username.Values) == 0 {
		return errors.BadRequest("go.micro.api.passenger", "Name cannot be blank")
	}
	password, ok := req.Get["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.passenger", "Pwd cannot be blank")
	}

	response, err := s.Client.Signup(ctx, &user.ReqSignup{
		Username: username.Values[0],
		Password: password.Values[0],
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
	username, ok := req.Get["username"]
	if !ok || len(username.Values) == 0 {
		return errors.BadRequest("go.micro.api.passenger", "Name cannot be blank")
	}
	password, ok := req.Get["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.passenger", "Pwd cannot be blank")
	}

	response, err := s.Client.Signin(ctx, &user.ReqSignin{
		Username: username.Values[0],
		Password: password.Values[0],
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

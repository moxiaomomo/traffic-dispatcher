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

// func (s *Say) Hello(ctx context.Context, req map[string]string, rsp *api.Response) error {
func (s *User) QueryUserByName(ctx context.Context, req *api.Request, rsp *api.Response) error {
	// logger.Info("Received User.QueryUserByName API request")
	// name, ok := req["name"]
	// if !ok || len(name) == 0 {
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.driver", "Name cannot be blank")
	}

	// 在restful api中调用rpc服务
	response, err := s.Client.QueryUserByName(ctx, &user.Request{
		UserID:   "testuserid",
		UserName: name.Values[0],
		UserPwd:  "somepwd",
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"user": response.GetUser(),
		"err":  response.GetError(),
		"msg":  response.GetSuccess(),
	})
	rsp.Body = string(b)

	return nil
}

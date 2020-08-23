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

func (s *User) Signup(ctx context.Context, req *api.Request, rsp *api.Response) error {
	reqUser, err := parseReqBody(req)
	if err != nil {
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
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
		return errors.BadRequest("go.micro.api.passenger", "request invalid")
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

package handler

import (
	"context"
	"encoding/json"
	"traffic-dispatcher/model"
	lbs "traffic-dispatcher/proto/lbs"

	api "github.com/micro/go-micro/v2/api/proto"
)

type Lbs struct {
	Client lbs.GeoLocationService
}

type GeoQueryBody struct {
	reqMsg model.WSMessage
}

// func parseReqOrderBody(req *api.Request) (reqGeo model.WSMessage, err error) {
// 	err = json.Unmarshal([]byte(req.Body), &reqGeo)
// 	return
// }

func (s *Lbs) QueryGeoNearby(ctx context.Context, req *api.Request, rsp *api.Response) error {
	// reqMsg, err := parseReqOrderBody(req)
	// if err != nil {
	// 	return errors.BadRequest("go.micro.api.passenger", "request invalid")
	// }
	response, err := s.Client.QueryGeoNearby(ctx, &lbs.QueryRequest{
		Data: []byte(req.Body),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.GetCode(),
		"data": response.GetData(),
		"msg":  response.GetMsg(),
	})
	rsp.Body = string(b)

	return nil
}

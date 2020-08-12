package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	api "github.com/micro/go-micro/v2/api/proto"
	// admin "path/to/service/proto/admin"
)

type Admin struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Admin.Call is called by the API as /admin/call with post body {"name": "foo"}
func (e *Admin) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Admin.Call request")

	// // extract the client from the context
	// adminClient, ok := client.AdminFromContext(ctx)
	// if !ok {
	// 	return errors.InternalServerError("go.micro.api.admin.admin.call", "admin client not found")
	// }

	// // TODO make request
	// response, err := adminClient.Call(ctx, &admin.Request{
	// 	Name: extractValue(req.Post["name"]),
	// })
	// if err != nil {
	// 	return errors.InternalServerError("go.micro.api.admin.admin.call", err.Error())
	// }

	// b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	// rsp.Body = string(b)
	rsp.Body = "Not Implemented"

	return nil
}

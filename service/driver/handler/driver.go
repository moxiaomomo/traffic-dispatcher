package handler

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/uber/h3-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"driver/client"
	"traffic-dispatcher/dbproxy"
	"traffic-dispatcher/model"
	driver "traffic-dispatcher/proto/driver"
	"traffic-dispatcher/proto/lbs"

	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
)

type Driver struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

func InsertGeo(resolution int, data model.WSMessage) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database("driverInfo").Collection("geoInfo")

	geo := h3.GeoCoord{
		Latitude:  data.Geo.Lat,
		Longitude: data.Geo.Lng,
	}
	h3Index := h3.FromGeo(geo, resolution)
	h3IndexStr := fmt.Sprintf("%#x", h3Index)

	doc := bson.M{
		"$set": model.Driver{
			Name:    data.User.Name,
			UID:     data.User.UID,
			H3Index: h3IndexStr,
			GeoInfo: bson.M{
				"type":        "Point",
				"coordinates": []float64{data.Geo.Lng, data.Geo.Lat},
			},
		},
	}

	opts := options.Update().SetUpsert(true)
	insertResult, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{"uid", data.User.UID}},
		doc,
		opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.UpsertedID)
}

// Driver.Call is called by the API as /driver/call with post body {"name": "foo"}
func (e *Driver) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Driver.Call request")

	// extract the client from the context
	driverClient, ok := client.DriverFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.driver.driver.call", "driver client not found")
	}

	// make request
	response, err := driverClient.Call(ctx, &lbs.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.driver.driver.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

type Greeter struct {
}

func (g *Greeter) HelloTest(ctx context.Context, req *driver.SayRequest, resp *driver.SayResponse) error {
	// fmt.Println("recived data")
	var data model.WSMessage
	if err := json.Unmarshal(req.Data, &data); err == nil {
		InsertGeo(7, data)
		resp.Greeting = "Hi " + req.Name
		return nil
	} else {
		resp.Greeting = "Oooops..." + req.Name
		return err
	}
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/uber/h3-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	log.Info("Inserted a single document: ", insertResult.UpsertedID)
}

func QueryGeo(lat float64, lng float64) (res []model.Driver, err error) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database("driverInfo").Collection("geoInfo")

	stages := mongo.Pipeline{}
	getNearbyStage := bson.D{
		{"$geoNear", bson.M{
			"near": bson.M{
				"type":        "Point",
				"coordinates": []float64{lng, lat},
			},
			"maxDistance":   100000,
			"spherical":     true,
			"distanceField": "distance",
		}}}

	stages = append(stages, getNearbyStage)

	filterCursor, err := collection.Aggregate(context.TODO(), stages)
	if err != nil {
		log.Error(err)
		return
	}
	for filterCursor.Next(context.TODO()) {
		var elem model.Driver
		err = filterCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return
		}
		res = append(res, elem)
	}
	return
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

func (g *Driver) ReportGeo(ctx context.Context, req *driver.ReportRequest, resp *driver.ReportResponse) error {
	var data model.WSMessage
	if err := json.Unmarshal(req.Data, &data); err == nil {
		InsertGeo(7, data)
		resp.Msg = "Hi " + req.Name
		return nil
	} else {
		resp.Msg = "Oooops..." + req.Name
		return err
	}
}

func (g *Driver) QueryGeo(ctx context.Context, req *driver.QueryRequest, resp *driver.QueryResponse) error {
	var data model.WSMessage
	var err error
	if err = json.Unmarshal(req.Data, &data); err == nil {
		if geolist, err := QueryGeo(data.Geo.Lat, data.Geo.Lng); err == nil {
			data, _ := json.Marshal(geolist)
			resp.Msg = "Hi " + req.Name
			resp.Data = data
			return nil
		}
	}
	resp.Msg = "Oooops..." + req.Name
	return err
}

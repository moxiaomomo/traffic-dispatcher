package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"traffic-dispatcher/dbproxy"
	"traffic-dispatcher/model"
	"traffic-dispatcher/proto/geo"

	"github.com/micro/go-micro/v2/util/log"
	h3 "github.com/uber/h3-go/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientDBMap = map[model.ClientRole]string{
	model.ClientDriver:    "driverInfo",
	model.ClientPassenger: "passengerInfo",
}

func InsertGeo(resolution int, data model.WSMessage) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database(clientDBMap[data.Role]).Collection("geoInfo")

	geo := h3.GeoCoord{
		Latitude:  data.Geo.Lat,
		Longitude: data.Geo.Lng,
	}
	h3Index := h3.FromGeo(geo, resolution)
	h3IndexStr := fmt.Sprintf("%#x", h3Index)

	doc := bson.M{
		"$set": model.UserLocation{
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

func QueryGeo(lat float64, lng float64, role model.ClientRole) (res []model.UserLocation, err error) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database(clientDBMap[role]).Collection("geoInfo")

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
		var elem model.UserLocation
		err = filterCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return
		}
		res = append(res, elem)
	}
	return
}

type GeoLocation struct{}

func (g *GeoLocation) ReportGeo(ctx context.Context, req *geo.ReportRequest, resp *geo.ReportResponse) error {
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

func (g *GeoLocation) QueryGeo(ctx context.Context, req *geo.QueryRequest, resp *geo.QueryResponse) error {
	var data model.WSMessage
	var err error
	if err = json.Unmarshal(req.Data, &data); err == nil {
		if geolist, err := QueryGeo(data.Geo.Lat, data.Geo.Lng, data.Role); err == nil {
			data, _ := json.Marshal(geolist)
			resp.Msg = "Hi " + req.Name
			resp.Data = data
			return nil
		}
	}
	resp.Msg = "Oooops..." + req.Name
	return err
}

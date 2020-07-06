package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"traffic-dispatcher/dbproxy"
	"traffic-dispatcher/model"

	h3 "github.com/uber/h3-go/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func testInsert(h3Index string, lat float64, lon float64) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database("driverInfo").Collection("geoInfo")

	d1 := model.Driver{
		Name:    "test1",
		UID:     "123456",
		H3Index: h3Index,
		// GeoType: "Point",
		// Coord:   []float64{lon, lat},
		GeoInfo: bson.M{
			"type":        "Point",
			"coordinates": []float64{lon, lat},
		},
	}
	// d1Json, err := json.Marshal(d1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	insertResult, err := collection.InsertOne(context.TODO(), d1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func testQuery(lat float64, lon float64) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database("driverInfo").Collection("geoInfo")

	stages := mongo.Pipeline{}
	getNearbyStage := bson.D{
		{"$geoNear", bson.M{
			"near": bson.M{
				"type":        "Point",
				"coordinates": []float64{lon, lat},
			},
			"maxDistance":   100000,
			"spherical":     true,
			"distanceField": "distance",
		}}}

	stages = append(stages, getNearbyStage)

	filterCursor, err := collection.Aggregate(context.TODO(), stages)
	if err != nil {
		log.Println(err)
	}
	for filterCursor.Next(context.TODO()) {
		var elem model.Driver
		err := filterCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v\n", elem)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	lat, err1 := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, err2 := strconv.ParseFloat(r.Form.Get("lon"), 64)
	if err1 != nil || err2 != nil {
		w.Write([]byte("FAILED"))
		return
	}

	geo := h3.GeoCoord{
		Latitude:  lat, // 37.775938728915946,
		Longitude: lon, // -122.41795063018799,
	}
	resolution := 9

	h3Index := h3.FromGeo(geo, resolution)
	h3IndexStr := fmt.Sprintf("%#x", h3Index)
	// Output:
	// 0x8928308280fffff

	neighbors := h3.KRing(h3Index, 1)
	for _, n := range neighbors {
		fmt.Printf("%#x\n", n)
	}

	// // test mongo
	// testInsert(h3IndexStr, lat, lon)
	testQuery(lat+0.1, lon+0.1)

	w.Write([]byte(h3IndexStr))
}

func main() {
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"traffic-dispatcher/dbproxy"
	"traffic-dispatcher/model"

	wsconn "traffic-dispatcher/connection"

	"github.com/gorilla/websocket"
	h3 "github.com/uber/h3-go/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func testInsert(resolution int, lat float64, lon float64) {
	dbCli := dbproxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database("driverInfo").Collection("geoInfo")

	geo := h3.GeoCoord{
		Latitude:  lat,
		Longitude: lon,
	}
	h3Index := h3.FromGeo(geo, resolution)
	h3IndexStr := fmt.Sprintf("%#x", h3Index)

	d1 := model.Driver{
		Name:    "test1",
		UID:     "123456",
		H3Index: h3IndexStr,
		GeoInfo: bson.M{
			"type":        "Point",
			"coordinates": []float64{lon, lat},
		},
	}

	insertResult, err := collection.InsertOne(context.TODO(), d1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func testQuery(lat float64, lon float64) (res []model.Driver, err error) {
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
		err = filterCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return
		}
		res = append(res, elem)
	}
	return
}

func testInsertHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	lat, err1 := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, err2 := strconv.ParseFloat(r.Form.Get("lon"), 64)
	if err1 != nil || err2 != nil {
		w.Write([]byte("FAILED"))
		return
	}

	resolution := 7

	for i := 1.0; i <= 10.0; i++ {
		var delta = 0.01 * i
		testInsert(resolution, lat+delta, lon+delta)
		testInsert(resolution, lat+delta, lon-delta)
		testInsert(resolution, lat-delta, lon+delta)
		testInsert(resolution, lat-delta, lon-delta)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := fmt.Sprintf(`{"code":0}`)
	w.Write([]byte(res))
}

func testQueryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()

	lat, err1 := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, err2 := strconv.ParseFloat(r.Form.Get("lon"), 64)
	if err1 != nil || err2 != nil {
		w.Write([]byte("FAILED"))
		return
	}

	// neighbors := h3.KRing(h3Index, 1)
	// for _, n := range neighbors {
	// 	fmt.Printf("%#x\n", n)
	// }

	// 经度0.1度约10km, 纬度0.1度约11.1km
	drivers, err := testQuery(lat+0.1, lon+0.1)
	if err != nil {
		res := fmt.Sprintf(`{"code":-1}`)
		w.Write([]byte(res))
		return
	}

	data, err := json.Marshal(drivers)
	if err != nil {
		res := fmt.Sprintf(`{"code":-2}`)
		w.Write([]byte(res))
		return
	}

	res := fmt.Sprintf(`{"code":0,"count":%d,"data":%s}`, len(drivers), string(data))
	w.Write([]byte(res))
}

func testWSHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *wsconn.WsConnection
	)

	// upgrade websocket
	if wsConn, err = wsconn.WsUpgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	// initiate connection
	if conn, err = wsconn.InitConnection(wsConn); err != nil {
		goto ERR
	}

	// 启动协程，持续发信息
	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte(`{"code":0}`)); err != nil {
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		if data, err := conn.ReadMessage(); err != nil {
			log.Println(string(data))
			goto ERR
		}
		// if err = conn.WriteMessage([]byte("ACK")); err != nil {
		// 	goto ERR
		// }
	}

ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/test/insert", testInsertHandler)
	http.HandleFunc("/test/query", testQueryHandler)
	http.HandleFunc("/ws", testWSHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

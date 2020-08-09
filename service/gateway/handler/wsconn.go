package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	wsconn "traffic-dispatcher/connection"
	"traffic-dispatcher/model"
	"traffic-dispatcher/proto/driver"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2"
)

var wsConnCount int

func reportGeoInfo(data []byte) {
	service := micro.NewService(
		micro.Name("client.service"),
	)
	service.Init()

	driverCli := driver.NewGreeterService("go.micro.api.driver", service.Client())
	rsp, err := driverCli.HelloTest(context.TODO(), &driver.SayRequest{
		Name: "Hello test",
		Data: data,
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(rsp.GetGreeting())
}

func WSConnHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *wsconn.WsConnection
		// 搜索范围的中心位置坐标
		wsMsg model.WSMessage
	)

	// 搜索附近坐标位置
	var processSearchLoc = func() {
		if wsMsg.Geo == (model.GeoLocation{}) || conn == nil {
			return
		}
		if drivers, err := QueryGeo(wsMsg.Geo.Lat, wsMsg.Geo.Lng); err == nil {
			if resp, err := json.Marshal(drivers); err == nil {
				conn.WriteMessage(resp)
			}
		}
	}

	// upgrade websocket
	if wsConn, err = wsconn.WsUpgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	// initiate connection
	if conn, err = wsconn.InitConnection(wsConn); err != nil {
		log.Println(err.Error())
		goto ERR
	}

	wsConnCount++
	log.Printf("Current connection count: %d\n", wsConnCount)

	// 启动协程，持续发信息
	go func() {
		for {
			processSearchLoc()
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		if data, err := conn.ReadMessage(); err != nil {
			log.Println(err.Error())
			goto ERR
		} else {
			if err := json.Unmarshal(data, &wsMsg); err == nil {
				if wsMsg.Command == model.CmdQueryGeo {
					processSearchLoc()
				} else if wsMsg.Command == model.CmdReportGeo {
					reportGeoInfo(data)
				}
			}
		}
	}

ERR:
	conn.Close()
	wsConnCount--
}

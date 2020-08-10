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

var (
	wsConn      *websocket.Conn
	conn        *wsconn.WsConnection
	wsConnCount int
	driverCli   driver.DriverSrvService
)

func init() {
	service := micro.NewService(
		micro.Name("client.service"),
	)
	service.Init()

	driverCli = driver.NewDriverSrvService("go.micro.api.driver", service.Client())
}

func reportGeoInfo(data []byte) {
	rsp, err := driverCli.ReportGeo(context.TODO(), &driver.ReportRequest{
		Name: "ReportGeoInfo",
		Data: data,
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(rsp.GetMsg())
}

// 搜索附近坐标位置
func queryGeoInfo(data []byte) {
	rsp, err := driverCli.QueryGeo(context.TODO(), &driver.QueryRequest{
		Name: "QueryGeoInfo",
		Data: data,
	})

	if err != nil {
		log.Println(err)
		return
	}

	// log.Println(rsp.GetData())
	conn.WriteMessage(rsp.GetData())
	// var geolist []model.Driver
	// if err := json.Unmarshal(rsp.GetData(), geolist); err == nil {
	// 	conn.WriteMessage(rsp.GetData())
	// }
}

// WSConnHandler websocket handler
func WSConnHandler(w http.ResponseWriter, r *http.Request) {
	// 搜索范围的中心位置坐标
	var wsMsg model.WSMessage
	var wsMsgByte []byte
	var err error

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
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println(wsMsg)
				if wsMsg.Geo == (model.GeoLocation{}) || conn == nil {
					// ...
				} else {
					queryGeoInfo(wsMsgByte)
				}
			}
		}
	}()

	for {
		if wsMsgByte, err = conn.ReadMessage(); err != nil {
			log.Println(err.Error())
			goto ERR
		} else {
			if err := json.Unmarshal(wsMsgByte, &wsMsg); err == nil {
				if wsMsg.Command == model.CmdQueryGeo {
					queryGeoInfo(wsMsgByte)
				} else if wsMsg.Command == model.CmdReportGeo {
					reportGeoInfo(wsMsgByte)
				}
			}
		}
	}

ERR:
	conn.Close()
	wsConnCount--
}

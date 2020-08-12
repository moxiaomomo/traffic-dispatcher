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
	"traffic-dispatcher/proto/passenger"
)

func reportGeoInfo(cliRole model.ClientRole, data []byte) {
	switch cliRole {
	case model.ClientDriver:
		if rsp, err := driverCli.ReportGeo(context.TODO(), &driver.ReportRequest{
			Name: "ReportGeoInfo",
			Data: data,
		}); err == nil {
			log.Println(rsp.GetMsg())
		} else {
			log.Println(err)
		}
		break
	case model.ClientPassenger:
		if rsp, err := passengerCli.ReportGeo(context.TODO(), &passenger.ReportPassengerRequest{
			Name: "ReportGeoInfo",
			Data: data,
		}); err == nil {
			log.Println(rsp.GetMsg())
		} else {
			log.Println(err)
		}
		break
	}
}

// 搜索附近坐标位置
func queryGeoInfo(cliRole model.ClientRole, data []byte) {
	switch cliRole {
	case model.ClientDriver:
		if rsp, err := driverCli.QueryGeo(context.TODO(), &driver.QueryRequest{
			Name: "QueryGeoInfo",
			Data: data,
		}); err == nil {
			conn.WriteMessage(rsp.GetData())
		} else {
			log.Println(err)
		}
		break
	case model.ClientPassenger:
		if rsp, err := passengerCli.QueryGeo(context.TODO(), &passenger.QueryPassengerRequest{
			Name: "QueryGeoInfo",
			Data: data,
		}); err == nil {
			conn.WriteMessage(rsp.GetData())
		} else {
			log.Println(err)
		}
		break
	}
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
				// log.Println(wsMsg)
				if wsMsg.Geo == (model.GeoLocation{}) || conn.IsClose() {
					// ...
				} else {
					queryGeoInfo(wsMsg.Role, wsMsgByte)
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
					queryGeoInfo(wsMsg.Role, wsMsgByte)
				} else if wsMsg.Command == model.CmdReportGeo {
					reportGeoInfo(wsMsg.Role, wsMsgByte)
				}
			}
		}
	}

ERR:
	conn.Close()

	wsMsg = model.WSMessage{}
	wsMsgByte = nil

	wsConnCount--
}

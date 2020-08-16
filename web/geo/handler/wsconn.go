package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"

	wsconn "traffic-dispatcher/connection"
	"traffic-dispatcher/model"
	geoProto "traffic-dispatcher/proto/geo"
)

// 上报坐标位置
func reportGeoInfo(cliRole model.ClientRole, data []byte) {

	if rsp, err := GeoCli.ReportGeo(context.TODO(), &geoProto.ReportRequest{
		Name: "ReportGeoInfo",
		Data: data,
	}); err == nil {
		log.Info(rsp.GetMsg())
	} else {
		log.Error(err.Error())
	}
}

// 搜索附近坐标位置
func queryGeoInfo(cliRole model.ClientRole, data []byte) {
	if rsp, err := GeoCli.QueryGeo(context.TODO(), &geoProto.QueryRequest{
		Name: "QueryGeoInfo",
		Data: data,
	}); err == nil {
		conn.WriteMessage(rsp.GetData())
	} else {
		log.Error(err.Error())
	}
}

// WSConnHandler websocket handler
func (g *GeoLocation) WSConnHandler(c *gin.Context) {
	// 搜索范围的中心位置坐标
	var wsMsg model.WSMessage
	var wsMsgByte []byte
	var err error

	// upgrade websocket
	if wsConn, err = wsconn.WsUpgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	// initiate connection
	if conn, err = wsconn.InitConnection(wsConn); err != nil {
		log.Info(err.Error())
		goto ERR
	}

	wsConnCount++
	log.Infof("Current connection count: %d\n", wsConnCount)

	// 启动协程，持续发信息
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// log.Info(wsMsg)
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
			log.Info(err.Error())
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

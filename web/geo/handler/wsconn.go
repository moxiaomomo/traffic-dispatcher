package handler

import (
	"context"
	"encoding/json"
	defultLog "log"
	"time"
	"traffic-dispatcher/util"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"

	"traffic-dispatcher/api/driver/mq"
	"traffic-dispatcher/config"
	"traffic-dispatcher/model"
	wsconn "traffic-dispatcher/net"
	lbsProto "traffic-dispatcher/proto/lbs"
)

// 上报坐标位置
func reportGeoInfo(cliRole model.ClientRole, data []byte) {

	if rsp, err := GeoCli.ReportGeo(context.TODO(), &lbsProto.ReportRequest{
		Name: "ReportGeoInfo",
		Data: data,
	}); err == nil {
		logger.Info(rsp.GetMsg())
	} else {
		logger.Error(err.Error())
	}
}

// 搜索附近坐标位置
func queryGeoInfo(param model.WSMessage) {
	data, _ := json.Marshal(param)
	// fmt.Printf("%+v\n", string(data))
	if rsp, err := GeoCli.QueryGeoNearby(context.TODO(), &lbsProto.QueryRequest{
		Name: "QueryGeoInfo",
		Data: data,
	}); err == nil {
		conn.WriteMessage(rsp.GetData())
	} else {
		logger.Error(err.Error())
	}
}

// WSConnHandler websocket handler
func (g *GeoLocation) WSConnHandler(c *gin.Context) {
	// 搜索范围的中心位置坐标
	var wsMsg model.WSMessage
	var wsMsgByte []byte
	var subMsg model.WSMessage
	var err error
	var roleStr string

	// upgrade websocket
	if wsConn, err = wsconn.WsUpgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	// initiate connection
	if conn, err = wsconn.InitConnection(wsConn); err != nil {
		logger.Info(err.Error())
		return
	}

	wsConnCount++
	// log.Infof("Current connection count: %d\n", wsConnCount) // not works
	defultLog.Printf("Current connection count: %v\n", wsConnCount)

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
				} else if subMsg.Command == model.CmdSubscribeGeo {
					// fmt.Printf("%+v\n", subMsg)
					queryGeoInfo(subMsg)
				}
			}
		}
	}()

	// 根据用户角色订阅不同topic信息
	roleStr = c.Query("role")
	if model.IsDriver(roleStr) {
		go mq.Subscribe(config.DriverLbsMQTopic)
	} else if model.IsPassenger(roleStr) {
		go mq.Subscribe(config.PassengerLbsMQTopic)
	}

	for {
		if wsMsgByte, err = conn.ReadMessage(); err != nil {
			logger.Info(err.Error())
			goto ERR
		} else {
			if err := json.Unmarshal(wsMsgByte, &wsMsg); err == nil {
				if wsMsg.Command == model.CmdQueryGeo {
					queryGeoInfo(wsMsg)
				} else if wsMsg.Command == model.CmdReportGeo {
					reportGeoInfo(wsMsg.Role, wsMsgByte)
				} else if wsMsg.Command == model.CmdSubscribeGeo {
					util.DeepCopyByGob(&subMsg, wsMsg)
				}
			}
		}
	}

ERR:
	conn.Close()
	logger.Info("===========connection closed===========")
	wsMsg = model.WSMessage{}
	wsMsgByte = nil

	wsConnCount--
}

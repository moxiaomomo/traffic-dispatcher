package handler

import (
	"context"
	"encoding/json"
	defultLog "log"
	"time"

	"traffic-dispatcher/config"
	"traffic-dispatcher/model"
	wsconn "traffic-dispatcher/net"
	wsnet "traffic-dispatcher/net"
	lbsProto "traffic-dispatcher/proto/lbs"
	"traffic-dispatcher/util"
	"traffic-dispatcher/web/geo/mq"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2/logger"
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
		conns[param.User.UID].WriteMessage(rsp.GetData())
	} else {
		logger.Error(err.Error())
	}
}

// WSConnHandler websocket handler
func (g *GeoLocation) WSConnHandler(c *gin.Context) {
	// 搜索范围的中心位置坐标
	var subMsg model.WSMessage
	var wsMsgByte []byte
	var err error
	var roleStr string

	// upgrade websocket
	var wsConn *websocket.Conn
	if wsConn, err = wsconn.WsUpgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	// initiate connection
	var conn *wsnet.WsConnection
	if conn, err = wsconn.InitConnection(wsConn); err != nil {
		logger.Info(err.Error())
		return
	}

	userID := c.Query("uid")
	if userID == "" {
		logger.Info("invalid user id")
		return
	}
	conns[userID] = conn

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
				if userInfos[userID] == nil {
					continue
				}
				if userInfos[userID].Geo == (model.GeoLocation{}) || conn.IsClose() {
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
		go mq.Subscribe(config.DriverLbsMQTopic, processSubscribeMessage)
	} else if model.IsPassenger(roleStr) {
		go mq.Subscribe(config.PassengerLbsMQTopic, processSubscribeMessage)
	}

	for {
		if wsMsgByte, err = conn.ReadMessage(); err != nil {
			logger.Info(err.Error())
			goto ERR
		} else {
			var wsMsg model.WSMessage
			if err := json.Unmarshal(wsMsgByte, &wsMsg); err == nil {
				userInfos[userID] = &wsMsg
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
	subMsg = model.WSMessage{}
	wsMsgByte = nil

	wsConnCount--
}

func processSubscribeMessage(topic string, msg string) error {
	for uid, conn := range conns {
		if userInfos[uid] == nil {
			continue
		}
		if topic == config.DriverLbsMQTopic && userInfos[uid].Role == model.ClientDriver {
			conn.WriteMessage([]byte(msg))
		} else if topic == config.PassengerLbsMQTopic && userInfos[uid].Role == model.ClientPassenger {
			conn.WriteMessage([]byte(msg))
		}
	}
	return nil
}

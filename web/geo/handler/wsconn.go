package handler

import (
	"context"
	"encoding/json"

	// defultLog "log"
	"time"

	"traffic-dispatcher/config"
	"traffic-dispatcher/model"
	wsconn "traffic-dispatcher/net"
	wsnet "traffic-dispatcher/net"
	lbsProto "traffic-dispatcher/proto/lbs"
	"traffic-dispatcher/util"
	"traffic-dispatcher/web/geo/mq"

	orderProto "traffic-dispatcher/proto/order"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2/logger"
	// "github.com/micro/go-micro/v2/broker"
)

type MsgResponse struct {
	Topic string      `json:"topic"`
	Data  interface{} `json:"data"`
}

func Init() {
	go mq.Subscribe(config.DriverLbsMQTopic, processSubscribeMessage)
	go mq.Subscribe(config.PassengerLbsMQTopic, processSubscribeMessage)
}

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

func writeMessage(wsConn *wsnet.WsConnection, data []byte) {
	if wsConn == nil {
		return
	}
	wsConn.WriteMessage(data)
}

// 搜索附近坐标位置
func queryGeoInfo(param model.WSMessage) {
	data, _ := json.Marshal(param)
	// fmt.Printf("%+v\n", string(data))
	if rsp, err := GeoCli.QueryGeoNearby(context.TODO(), &lbsProto.QueryRequest{
		Name: "QueryGeoInfo",
		Data: data,
	}); err == nil {
		resp := MsgResponse{
			Topic: "geolist",
			Data:  string(rsp.GetData()),
		}
		// fmt.Printf("resp: %+v\n", resp)
		respB, _ := json.Marshal(resp)
		// fmt.Printf("respB: %+v\n", respB)
		writeMessage(conns[param.User.UID], respB)
	} else {
		logger.Error(err.Error())
	}
}

// 获取订单信息
func queryOrderHis(userID string, role string) {
	if rsp, err := OrderCli.QueryOrderHis(context.TODO(), &orderProto.ReqOrderHis{
		UserId: userID,
		Role:   int32(model.RoleValue(role)),
	}); err == nil {
		resp := MsgResponse{
			Topic: "orderhis",
			Data:  rsp.GetOrders(),
		}
		respB, _ := json.Marshal(resp)
		// logger.Infof("%t %s", conns[userID] == nil, string(respB))
		writeMessage(conns[userID], respB)
	} else {
		logger.Error(err.Error())
	}
}

// WSConnHandler websocket handler
func (g *GeoLocation) WSConnHandler(c *gin.Context) {
	// 搜索范围的中心位置坐标
	var wsMsgByte []byte
	var err error
	var roleStr string
	var userID string

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

	userID = c.Query("uid")
	if userID == "" {
		logger.Info("invalid user id")
		return
	}
	conns[userID] = conn

	wsConnCount++
	logger.Infof("Current connection count: %d\n", wsConnCount) // not works?
	// defultLog.Printf("Current connection count: %v\n", wsConnCount)

	// 启动协程，持续发信息
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// logger.Infof("map len: %d %d\n", len(userInfos), len(subInfos))
				if userInfos[userID] == nil {
					continue
				}
				if userInfos[userID].Geo == (model.GeoLocation{}) || conn.IsClose() {
					// ...
				}
				if subInfos[userID] != nil {
					// logger.Infof("tick task: %+v\n", subInfos)
					queryGeoInfo(*subInfos[userID])
					queryOrderHis(userID, roleStr)
				}
			}
		}
	}()

	// 根据用户角色订阅不同topic信息
	roleStr = c.Query("role")
	userID = c.Query("uid")

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
					var subMsg model.WSMessage
					util.DeepCopyByGob(&subMsg, wsMsg)
					subInfos[userID] = &subMsg
				}
			}
		}
	}

ERR:
	conn.Close()
	logger.Info("===========connection closed===========")
	wsMsgByte = nil

	wsConnCount--
}

func processSubscribeMessage(topic string, msg string) error {
	var order orderProto.Order
	err := json.Unmarshal([]byte(msg), &order)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

//	// geo parse
//	srcGeo := util.ParseGeoLocation(order.SrcGeo)
//	srcAddr, _ := geoCoder.ReverseGeocode(srcGeo.Lat, srcGeo.Lng)
//        logger.Infof("%+v %+v %+v\n", order.SrcGeo, srcGeo, srcAddr)
//	if srcAddr != nil {
//		order.SrcAddr = srcAddr.FormattedAddress
//	}
//	destGeo := util.ParseGeoLocation(order.DestGeo)
//	destAddr, _ := geoCoder.ReverseGeocode(destGeo.Lat, destGeo.Lng)
//	if destAddr != nil {
//		order.DestAddr = destAddr.FormattedAddress
//	}

	orderJSON, _ := json.Marshal(order)
	resp := MsgResponse{
		Topic: "orderreq",
		Data:  string(orderJSON),
	}
	logger.Infof("on subscribe message: %+v\n", resp)
	// logger.Infof("current conn map: %+v\n", userInfos)
	respB, _ := json.Marshal(resp)

	for uid, conn := range conns {
		if userInfos[uid] == nil {
			continue
		}
		logger.Infof("topic: %s, role: %d\n", topic, userInfos[uid].Role)
		if topic == config.DriverLbsMQTopic && subInfos[uid].Role == model.ClientDriver {
			logger.Infof("%b, to send messageto driver", conn == nil)
			writeMessage(conn, respB)
		} else if topic == config.PassengerLbsMQTopic && subInfos[uid].Role == model.ClientPassenger {
			writeMessage(conn, respB)
		}
	}
	return nil
}

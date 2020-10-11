package handler

import (
	wsnet "traffic-dispatcher/net"
	"traffic-dispatcher/proto/lbs"
	"traffic-dispatcher/proto/order"

	"github.com/gorilla/websocket"
)

var (
	wsConn      *websocket.Conn
	conn        *wsnet.WsConnection
	wsConnCount int
	GeoCli      lbs.GeoLocationService
	OrderCli    order.OrderService
)

type GeoLocation struct {
}

func init() {
}

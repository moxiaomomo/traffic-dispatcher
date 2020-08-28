package handler

import (
	wsnet "traffic-dispatcher/net"
	"traffic-dispatcher/proto/lbs"

	"github.com/gorilla/websocket"
)

var (
	wsConn      *websocket.Conn
	conn        *wsnet.WsConnection
	wsConnCount int
	GeoCli      lbs.GeoLocationService
)

type GeoLocation struct {
}

func init() {
}

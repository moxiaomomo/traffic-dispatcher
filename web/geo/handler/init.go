package handler

import (
	wsconn "traffic-dispatcher/connection"
	"traffic-dispatcher/proto/geo"

	"github.com/gorilla/websocket"
)

var (
	wsConn      *websocket.Conn
	conn        *wsconn.WsConnection
	wsConnCount int
	GeoCli      geo.GeoLocationService
)

type GeoLocation struct {
}

func init() {
}

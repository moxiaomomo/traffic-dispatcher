package handler

import (
	"traffic-dispatcher/model"
	wsnet "traffic-dispatcher/net"
	"traffic-dispatcher/proto/lbs"
)

var (
	conns       map[string]*wsnet.WsConnection
	userInfos   map[string]*model.WSMessage
	wsConnCount int
	GeoCli      lbs.GeoLocationService
)

type GeoLocation struct {
}

func init() {
}

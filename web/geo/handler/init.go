package handler

import (
	"traffic-dispatcher/model"
	wsnet "traffic-dispatcher/net"
	"traffic-dispatcher/proto/lbs"
	"traffic-dispatcher/proto/order"
)

var (
	conns       map[string]*wsnet.WsConnection
	userInfos   map[string]*model.WSMessage
	subInfos    map[string]*model.WSMessage
	wsConnCount int
	GeoCli      lbs.GeoLocationService
	OrderCli    order.OrderService
)

type GeoLocation struct {
}

func init() {
}

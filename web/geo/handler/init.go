package handler

import (
	"traffic-dispatcher/model"
	wsnet "traffic-dispatcher/net"
	"traffic-dispatcher/proto/lbs"
	"traffic-dispatcher/proto/order"

	geoGo "github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

var (
	conns       map[string]*wsnet.WsConnection
	userInfos   map[string]*model.WSMessage
	subInfos    map[string]*model.WSMessage
	wsConnCount int
	GeoCli      lbs.GeoLocationService
	OrderCli    order.OrderService

	geoCoder geoGo.Geocoder
)

type GeoLocation struct {
}

func init() {
	conns = make(map[string]*wsnet.WsConnection)
	userInfos = make(map[string]*model.WSMessage)
	subInfos = make(map[string]*model.WSMessage)
	geoCoder = openstreetmap.Geocoder()
}

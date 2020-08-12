package handler

import (
	wsconn "traffic-dispatcher/connection"
	"traffic-dispatcher/proto/admin"
	"traffic-dispatcher/proto/driver"
	"traffic-dispatcher/proto/passenger"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2"
)

var (
	wsConn       *websocket.Conn
	conn         *wsconn.WsConnection
	wsConnCount  int
	driverCli    driver.DriverSrvService
	passengerCli passenger.PassengerSrvService
	adminCli     admin.AdminService
)

func init() {
	drvSvc := micro.NewService(micro.Name("driver.client"))
	drvSvc.Init()
	driverCli = driver.NewDriverSrvService("go.micro.api.driver", drvSvc.Client())

	psgSvc := micro.NewService(micro.Name("passenger.client"))
	psgSvc.Init()
	passengerCli = passenger.NewPassengerSrvService("go.micro.api.passenger", psgSvc.Client())

	adminSvc := micro.NewService(micro.Name("admin.client"))
	adminSvc.Init()
	adminCli = admin.NewAdminService("go.micro.api.admin", adminSvc.Client())
}

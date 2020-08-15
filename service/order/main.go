package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	hello "traffic-dispatcher/proto/hello"
	"traffic-dispatcher/service/order/handler"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterSayHandler(service.Server(), new(handler.Order))

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

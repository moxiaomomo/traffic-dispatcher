package main

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"

	"driver/client"
	"driver/handler"
	"driver/notification"

	"github.com/micro/go-micro/v2"

	driver "driver/proto/driver"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.driver"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Driver service client
		micro.WrapHandler(client.DriverWrapper(service)),
	)

	// Register Handler
	driver.RegisterDriverHandler(service.Server(), new(handler.Driver))

	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error :%v\n", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error:%v\n", err)
	}
	go notification.Subscribe("lbs.dispatcher.task")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

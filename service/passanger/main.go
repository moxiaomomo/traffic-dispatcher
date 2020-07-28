package main

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"

	"passanger/client"
	"passanger/handler"
	"passanger/notification"

	"github.com/micro/go-micro/v2"

	passanger "passanger/proto/passanger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.passanger"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Passanger service client
		micro.WrapHandler(client.PassangerWrapper(service)),
	)

	// Register Handler
	passanger.RegisterPassangerHandler(service.Server(), new(handler.Passanger))

	log.Info(broker.String())
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error :%v\n", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error:%v\n", err)
	}

	go notification.Publish("lbs.dispatcher.task")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

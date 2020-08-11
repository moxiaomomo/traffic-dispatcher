package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"

	"passenger/client"
	"passenger/handler"
	"passenger/notification"
	passenger "traffic-dispatcher/proto/passenger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.passenger"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Passenger service client
		micro.WrapHandler(client.PassengerWrapper(service)),
	)

	// Register Handler
	passenger.RegisterPassengerSrvHandler(service.Server(), new(handler.Passenger))

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

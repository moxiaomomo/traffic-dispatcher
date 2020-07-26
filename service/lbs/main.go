package main

import (
	"lbs/handler"
	"lbs/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"traffic-dispatcher/proto/lbs"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.lbs"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	lbs.RegisterLbsHandler(service.Server(), new(handler.Lbs))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.lbs", service.Server(), new(subscriber.Lbs))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

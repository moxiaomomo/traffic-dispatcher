package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"user/api/handler"
	"user/api/subscriber"

	api "user/api/proto/api"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.api"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	api.RegisterApiHandler(service.Server(), new(handler.Api))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.api", service.Server(), new(subscriber.Api))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

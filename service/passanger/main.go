package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"passanger/handler"
	"passanger/client"

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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

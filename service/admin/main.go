package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"admin/handler"
	"admin/client"

	admin "admin/proto/admin"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.admin"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Admin service client
		micro.WrapHandler(client.AdminWrapper(service)),
	)

	// Register Handler
	admin.RegisterAdminHandler(service.Server(), new(handler.Admin))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

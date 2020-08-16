package main

import (
	"log"
	"traffic-dispatcher/proto/geo"
	"traffic-dispatcher/web/geo/handler"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.web.geo"),
	)
	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	handler.GeoCli = geo.NewGeoLocationService("go.micro.srv.lbs", client.DefaultClient)

	// Create RESTful handler (using Gin)
	geoLoc := new(handler.GeoLocation)
	router := gin.Default()
	router.GET("/ws/lbs", geoLoc.WSConnHandler)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

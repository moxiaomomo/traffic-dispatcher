package main

import (
	"gateway/handler"
	"log"

	"github.com/micro/go-micro/v2/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
		web.Address(":8082"),
	)
	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	service.HandleFunc("/", handler.ApiHandler)

	service.HandleFunc("/hello", handler.HelloHandler)
	service.HandleFunc("/test/insert", handler.InsertGeoHandler)
	service.HandleFunc("/test/query", handler.QueryGeoHandler)
	service.HandleFunc("/ws/lbs", handler.WSConnHandler)
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

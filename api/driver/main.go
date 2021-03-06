package main

import (
	"log"
	"traffic-dispatcher/api/driver/handler"
	"traffic-dispatcher/api/driver/mq"
	"traffic-dispatcher/config"
	"traffic-dispatcher/proto/order"
	user "traffic-dispatcher/proto/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
)

// 将grpc服务转为restful接口
// 微服务是多个独立的服务，api网关提供单个入口，整合微服务提供统一的api
func main() {
	service := micro.NewService(
		// go.micro.api是默认命名空间，访问api需要带上user命名空间，如：/user/xx
		// 如果不想使用默认命名空间可以在启动服务是设置
		micro.Name("go.micro.api.driver"),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.User{Client: user.NewUserService("go.micro.srv.user", service.Client())},
		),
	)

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Order{Client: order.NewOrderService("go.micro.srv.order", service.Client())},
		),
	)

	// init broker
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error :%v\n", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error:%v\n", err)
	}
	go mq.Subscribe(config.DriverLbsMQTopic)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

package main

import (
	"traffic-dispatcher/api/driver/handler"
	hello "traffic-dispatcher/proto/hello"

	"github.com/micro/go-micro/v2"
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
			&handler.Say{Client: hello.NewSayService("go.micro.srv.order", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

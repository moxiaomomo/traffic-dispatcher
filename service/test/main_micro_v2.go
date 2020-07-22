package main

import (
	"context"
	"fmt"

	hello "traffic-dispatcher/proto/hello"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

type SayV2 struct{}

func (s *SayV2) Hello(ctx context.Context, req *hello.SayRequest, rsp *hello.SayResponse) error {
	fmt.Printf("received req %#v \n", req)
	rsp.From = "server"
	rsp.To = "client"
	rsp.Msg = "ok"
	return nil
}

// test gomicro v2
func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://192.168.2.244:2379"}
	})

	service := micro.NewService(
		micro.Name("xiaomo.srv.say"),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	service.Init()

	hello.RegisterSayHandler(service.Server(), new(SayV2))

	if err := service.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"fmt"

	"traffic-dispatcher/proto/hello"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.SayRequest, rsp *hello.SayResponse) error {
	fmt.Printf("received req %#v \n", req)
	rsp.From = "server"
	rsp.To = "client"
	rsp.Msg = "ok"
	return nil
}

// test gomicro v1
func main_v1() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://192.168.2.244:2379"}
	})

	service := micro.NewService(
		micro.Name("xiaomo.srv.say"),
		micro.Registry(reg),
	)
	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		panic(err)
	}
}

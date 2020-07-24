package main

import (
    "context"
    "fmt"
    proto "traffic-dispatcher/service/admin/proto"

    "github.com/micro/go-micro/v2"
)

type LBSSrv struct{}

func (g *LBSSrv) DriverLocations(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
    rsp.Msg = "Name: " + req.Name // 对客户端传递的字符串做处理
    return nil
}

func main() {
    // 创建服务器
    service := micro.NewService(
        micro.Name("lbssrv"),
    )
    service.Init()
    // 注册 handler
    proto.RegisterLBSSrvHandler(service.Server(), new(LBSSrv))

    if err := service.Run(); err != nil {
        fmt.Println(err.Error())
    }
}
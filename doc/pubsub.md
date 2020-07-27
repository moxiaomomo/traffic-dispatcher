- driver api 服务订阅

```shell
cd traffic-dispatcher/service/driver
make build
./driver-api --registry=etcd --registry_address=192.168.2.244:2379
```

- passanger api 服务订阅

```shell
cd traffic-dispatcher/service/passanger
make build
./passanger-api --registry=etcd --registry_address=192.168.2.244:2379
```

- broker 初始化示例

```go
package main

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"

	"driver/client"
	"driver/handler"
	"driver/notification"

	"github.com/micro/go-micro/v2"

	driver "driver/proto/driver"
)

func main() {
    // ...

	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error :%v\n", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error:%v\n", err)
    }
    // 订阅
    go notification.Subscribe("test.topic")

    // 发布
    // go notification.Publish("test.topic")

    // ...
}
```

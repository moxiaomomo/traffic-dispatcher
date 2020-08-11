- driver api 服务订阅

```shell
cd traffic-dispatcher/service/driver
make build
./driver-api --registry=etcd --registry_address=192.168.2.244:2379
```

- passenger api 服务订阅

```shell
cd traffic-dispatcher/service/passenger
make build
./passenger-api --registry=etcd --registry_address=192.168.2.244:2379
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

- broker 用 rabbitmq 替代 http

  - 启动 rabbitmq 服务

```shell
# /www/rabbitmq目录可自定义，主要用于目录挂载
$mkdir -p /www/rabbitmq
$docker run -d --hostname rabbit-node1 --name rabbit-node1 -p 5672:5672 -p15672:15672 -v /www/rabbitmq:/var/lib/rabbitmq rabbitmq:management
```

- 先在 main 包中导入 rabbitmq 插件包

```go
package main

import (
	"github.com/micro/go-micro/v2/broker"
	// ...
)
```

- 启动时传入指定参数

```shell
// 司机服务
./driver-api --registry=etcd --registry_address=192.168.2.244:2379 --broker=rabbitmq --broker_address=amqp://guest:guest@192.168.2.244:5672
// 乘客服务
./passenger-api --registry=etcd --registry_address=192.168.2.244:2379 --broker=rabbitmq --broker_address=amqp://guest:guest@192.168.2.244:5672
```

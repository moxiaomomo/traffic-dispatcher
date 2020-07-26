- `micro server`开启后日志报错:

```
2020-07-24 20:51:45  file=manager/events.go:124 level=warn service=runtime Error processing create event for service helloworld:latest in namespace go.micro: error: 路径规格 'master' 未匹配任何 git 已知文件
(micro server命令不能在micro包所在的目录中运行...切换到其他目录运行则不再报错)


2020-07-24 21:01:02 file=handler/handler.go:227 level=error service=debug Error calling micro/helloworld@192.168.2.244:37145 ({"id":"go.micro.client","code":500,"detail":"malformed method name: \"/micro/helloworld.Debug/Trace\"","status":"Internal Server Error"})
```

- `micro api`无法访问 etcd

```
xiaomo@xiaomo:~$ micro --registry=etcd --registry_address=192.168.2.244:2379 api --handler=api
{"id":"go.micro.client","code":500,"detail":"connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:8081: connect: connection refused\"","status":"Internal Server Error"}

// micro --registry=etcd --registry_address=192.168.2.244:2379 api  (无法启动)
micro --registry=etcd --registry_address=192.168.2.244:2379 (正常启动)
```

- 导入本地包的问题

```
xiaomo@xiaomo:/data/go/src/traffic-dispatcher/service/passanger$ make build
protoc --proto_path=. --micro_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. --go_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. proto/passanger/passanger.proto
go build -o passanger-api *.go
../lbs/main.go:6:2: package lbs/handler is not in GOROOT (/usr/local/go/src/lbs/handler)
client/passanger.go:6:2: import "lbs/proto" is a program, not an importable package
../lbs/main.go:9:2: module lbs/proto@latest found (v0.0.0-00010101000000-000000000000, replaced by ../lbs), but does not contain package lbs/proto/lbs
../lbs/main.go:7:2: package lbs/subscriber is not in GOROOT (/usr/local/go/src/lbs/subscriber)
make: *** [Makefile:14：build] 错误 1


client/passanger.go:6:2: module traffic-dispatcher/proto/lbs@latest found (v0.0.0-00010101000000-000000000000, replaced by ../../proto), but does not contain package traffic-dispatcher/proto/lbs

// https://www.cnblogs.com/t0000/p/13354257.html
// 参考 test/test_pkg的示例代码
```

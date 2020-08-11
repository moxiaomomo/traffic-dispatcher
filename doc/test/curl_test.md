- 部署 lbs service

```shell
$ cd traffic-dispatcher/service/lbs
$ make build
$ ./lbs-service --registry=etcd --registry_address=xx.xx.xx.xx:2379
```

日志输出

```
2020-07-26 22:32:08  file=v2@v2.9.1/service.go:200 level=info Starting [service] go.micro.service.lbs
2020-07-26 22:32:08  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:45639
2020-07-26 22:32:08  file=grpc/grpc.go:881 level=info Broker [http] Connected to 127.0.0.1:34061
2020-07-26 22:32:08  file=grpc/grpc.go:697 level=info Registry [etcd] Registering node: go.micro.service.lbs-6e414f79-bb9d-441b-8a65-e35638d5ff87
2020-07-26 22:32:08  file=grpc/grpc.go:730 level=info Subscribing to topic: go.micro.service.lbs

2020-07-26 22:34:03  file=handler/lbs.go:14 level=info Received Lbs.Call request

2020-07-26 22:39:22  file=handler/lbs.go:14 level=info Received Lbs.Call request
```

- 部署 passange api

```shell
$ cd traffic-dispatcher/service/passenger
$ make build
$ ./passenger-api --registry=etcd --registry_address=xx.xx.xx.xx:2379
```

日志输出

```
2020-07-26 22:39:02  file=v2@v2.9.1/service.go:200 level=info Starting [service] go.micro.api.passenger
2020-07-26 22:39:02  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:43693
2020-07-26 22:39:02  file=grpc/grpc.go:697 level=info Registry [etcd] Registering node: go.micro.api.passenger-8fdf18b3-6a04-4adc-94eb-d7bb5280675a

2020-07-26 22:39:22  file=handler/passenger.go:30 level=info Received Passenger.Call request
2020-07-26 22:40:16  file=handler/passenger.go:56 level=info Received Passenger.GetLocation request
```

- 部署 micro api 网关

```shell
$ micro --registry=etcd --registry_address=xx.xx.xx.xx:2379 api
```

日志输出

```
2020-07-26 22:38:48  file=api/api.go:285 level=info service=api Registering API Default Handler at /
2020-07-26 22:38:48  file=http/http.go:90 level=info service=api HTTP API Listening on [::]:8080
2020-07-26 22:38:48  file=v2@v2.9.1/service.go:200 level=info service=api Starting [service] go.micro.api
2020-07-26 22:38:48  file=grpc/grpc.go:864 level=info service=api Server [grpc] Listening on [::]:37859
2020-07-26 22:38:48  file=grpc/grpc.go:697 level=info service=api Registry [etcd] Registering node: go.micro.api-2ad4ed01-6756-4919-932d-c270fdfa961a
2020-07-26 22:38:50  file=registry/registry.go:102 level=error service=api unable to get service: service not found
127.0.0.1 - - [26/Jul/2020:22:39:12 +0800] "GET /passenger/getlocation HTTP/1.1" 500 117 "" "curl/7.68.0"
127.0.0.1 - - [26/Jul/2020:22:39:22 +0800] "GET /passenger/call HTTP/1.1" 200 48 "" "curl/7.68.0"
127.0.0.1 - - [26/Jul/2020:22:40:16 +0800] "GET /passenger/getLocation HTTP/1.1" 200 45 "" "curl/7.68.0"
```

- curl 测试

```
$ curl http://localhost:8080/passenger/call
{"statusCode":200,"body":"{\"msg\":\"Hello \"}"}
$ curl http://localhost:8080/passenger/getLocation
{"statusCode":200,"body":"{\"code\": \"0\"}"}
```

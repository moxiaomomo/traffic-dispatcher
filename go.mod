module traffic-dispatcher

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1 // indirect
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/uber/h3-go/v3 v3.0.2 // indirect
	go.mongodb.org/mongo-driver v1.3.4
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	traffic-dispatcher/proto v1.0.0
)

replace traffic-dispatcher/proto => ./proto

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

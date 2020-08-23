module traffic-dispatcher

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1 // indirect
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/uber/h3-go/v3 v3.0.2
	go.mongodb.org/mongo-driver v1.4.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	traffic-dispatcher/config v1.0.0
	traffic-dispatcher/net v1.0.0
	traffic-dispatcher/db v1.0.0
	traffic-dispatcher/model v1.0.0
	traffic-dispatcher/proto v1.0.0
	traffic-dispatcher/util v1.0.0
)

replace traffic-dispatcher/proto => ./proto

replace traffic-dispatcher/config => ./config

replace traffic-dispatcher/model => ./model

replace traffic-dispatcher/db => ./db

replace traffic-dispatcher/net => ./net

replace traffic-dispatcher/util => ./util

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

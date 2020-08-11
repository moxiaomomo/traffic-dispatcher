module gateway

go 1.14

require (
	github.com/gorilla/websocket v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/uber/h3-go/v3 v3.0.2
	go.mongodb.org/mongo-driver v1.4.0
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	traffic-dispatcher/connection v0.0.0-00010101000000-000000000000
	traffic-dispatcher/dbproxy v0.0.0-00010101000000-000000000000
	traffic-dispatcher/model v0.0.0-00010101000000-000000000000
	traffic-dispatcher/proto v1.0.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace traffic-dispatcher/connection => ../../connection

replace traffic-dispatcher/dbproxy => ../../dbproxy

replace traffic-dispatcher/model => ../../model

replace traffic-dispatcher/proto => ../../proto

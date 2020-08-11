module driver

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1
	github.com/uber/h3-go v3.0.1+incompatible
	go.mongodb.org/mongo-driver v1.4.0
	google.golang.org/protobuf v1.25.0
	traffic-dispatcher/proto v0.0.0-00010101000000-000000000000
	traffic-dispatcher/connection v0.0.0-00010101000000-000000000000
	traffic-dispatcher/dbproxy v0.0.0-00010101000000-000000000000
	traffic-dispatcher/model v0.0.0-00010101000000-000000000000
)

replace traffic-dispatcher/proto => ../../proto

replace traffic-dispatcher/connection => ../../connection

replace traffic-dispatcher/dbproxy => ../../dbproxy

replace traffic-dispatcher/model => ../../model


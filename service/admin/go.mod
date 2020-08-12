module admin

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
    github.com/micro/go-micro/v2 v2.9.1
    traffic-dispatcher/proto v1.0.0
)

replace traffic-dispatcher/proto => ../../proto

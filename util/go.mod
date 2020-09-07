module util

go 1.14

require (
	github.com/micro/go-micro/v2 v2.9.1
	traffic-dispatcher/model v1.0.0
	traffic-dispatcher/proto v1.0.0
)

replace traffic-dispatcher/model => ../model

replace traffic-dispatcher/proto => ../proto

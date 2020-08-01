module gateway

go 1.14

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/uber/h3-go/v3 v3.0.2 // indirect
	go.mongodb.org/mongo-driver v1.4.0 // indirect
	traffic-dispatcher/connection v0.0.0-00010101000000-000000000000 // indirect
	traffic-dispatcher/dbproxy v0.0.0-00010101000000-000000000000 // indirect
	traffic-dispatcher/model v0.0.0-00010101000000-000000000000 // indirect
)

replace traffic-dispatcher/connection => ../../connection

replace traffic-dispatcher/dbproxy => ../../dbproxy

replace traffic-dispatcher/model => ../../model

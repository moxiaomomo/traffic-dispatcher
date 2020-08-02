module bot-driver

go 1.14

require (
	go.mongodb.org/mongo-driver v1.4.0 // indirect
	traffic-dispatcher/model v0.0.0-00010101000000-000000000000
)

replace traffic-dispatcher/model => ../../../model

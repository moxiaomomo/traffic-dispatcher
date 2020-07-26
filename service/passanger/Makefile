
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/passanger/passanger.proto
    

.PHONY: build
build: proto

	go build -o passanger-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t passanger-api:latest

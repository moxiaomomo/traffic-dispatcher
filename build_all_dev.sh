#!/bin/bash
go build -o bin/svc_user service/user/main.go
go build -o bin/svc_order service/order/main.go
go build -o bin/svc_lbs service/lbs/main.go
go build -o bin/api_driver api/driver/main.go
go build -o bin/api_passenger api/passenger/main.go
go build -o bin/web_geo web/geo/main.go
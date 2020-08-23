module db

go 1.14

require (
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v1.8.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.4.0
	traffic-dispatcher/config v1.0.0
	traffic-dispatcher/model v1.0.0
)

replace traffic-dispatcher/model => ../model

replace traffic-dispatcher/config => ../config

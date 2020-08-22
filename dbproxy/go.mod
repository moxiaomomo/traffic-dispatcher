module dbproxy

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	go.mongodb.org/mongo-driver v1.4.0
	traffic-dispatcher/model v1.0.0
)

replace traffic-dispatcher/model => ../model

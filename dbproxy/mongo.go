package dbproxy

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

// MongoConn : create mongo connection
func MongoConn() *mongo.Client {
	var err error
	if mongoCli == nil {
		// 设置客户端连接配置
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		// 连接到MongoDB
		mongoCli, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("create mongo client err: " + err.Error())
		}
	}

	// 检查连接
	err = mongoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("ping mongo client err: " + err.Error())
	}

	return mongoCli
}

package mongo

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"traffic-dispatcher/config"
)

var mongoCli *mongo.Client

// MongoConn : create mongo connection
func MongoConn() *mongo.Client {
	var err error
	if mongoCli == nil {
		// 设置客户端连接配置
		clientOptions := options.Client().ApplyURI(config.MongoSource)
		// 连接到MongoDB
		mongoCli, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			logger.Fatal("create mongo client err: " + err.Error())
		}
	}

	// 检查连接
	err = mongoCli.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("ping mongo client err: " + err.Error())
	}

	return mongoCli
}

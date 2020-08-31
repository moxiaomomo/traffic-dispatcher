package mq

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

// Publish : 创建一个发布者
func Publish(topic string, msg string) {
	// 创建一个消息
	mqMsg := &broker.Message{
		Header: map[string]string{
			"from": "passenger",
		},
		Body: []byte(msg),
	}
	// 打印 broker
	log.Info(broker.String())
	// 发布消息
	if err := broker.Publish(topic, mqMsg); err != nil {
		log.Info("[pub] Message publication failed: ", err)
	} else {
		log.Info("[pub] Message published: ", string(mqMsg.Body))
	}
}

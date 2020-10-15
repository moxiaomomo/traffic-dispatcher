package mq

import (
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

// Message 消息体
type Message struct {
	Header map[string]string
	Body   []byte
}

// Subscribe 订阅topic
func Subscribe(topic string, callback func(topic string, data string) error) {
	// 订阅消息
	log.Infof("[sub]Ready to subscribe message, topic:%s\n", topic)
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		callback(topic, string(p.Message().Body))
		return nil
	})

	if err != nil {
		log.Error(err.Error())
	}
}

package notification

import (
	"fmt"

	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

// Message 消息体
type Message struct {
	Header map[string]string
	Body   []byte
}

// Subscribe 订阅topic
func Subscribe(topic string) {
	// 订阅消息
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

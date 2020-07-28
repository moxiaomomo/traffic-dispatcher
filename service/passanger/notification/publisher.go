package notification

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

// var topic = "go.micro.msg.order"

// Publish : 创建一个发布者, 并每秒钟给主题发送一次信息
func Publish(topic string) {
	log.Infof("[pub]Ready to publish message, topic:%s\n", topic)
	// 创建一个每秒钟执行的定时器
	tick := time.NewTicker(time.Second * 3)
	i := 0
	// 定时器开始执行
	for range tick.C {
		// 创建一个消息
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d:%s", i, time.Now().String())),
		}

		// 发布消息
		if err := broker.Publish(topic, msg); err != nil {
			log.Infof("[pub] Message publication failed: %+v\n", err)
		} else {
			log.Infof("[pub] Message published suc: %s\n", string(msg.Body))
		}
		i++
	}
}

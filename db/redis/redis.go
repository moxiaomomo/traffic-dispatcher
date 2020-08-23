package redis

import (
	"fmt"
	"time"

	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/logger"

	"traffic-dispatcher/config"
)

var (
	pool *redis.Pool
)

// newRedisPool : 创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 1. 打开连接
			c, err := redis.Dial("tcp", config.RedisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// 2. 访问认证
			if _, err = c.Do("AUTH", config.RedisPass); err != nil {
				fmt.Println(err)
				c.Close()
				return nil, err
			}
			return c, nil
		},
		// 检查连接的有效性
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
	_, err := pool.Get().Do("KEYS", "*")
	if err != nil {
		logger.Errorf("Failed to initiate redis pool, %v", err)
	} else {
		logger.Infof("Success to initate redis pool")
	}
}

// ConnPool redis连接池
func ConnPool() *redis.Pool {
	return pool
}

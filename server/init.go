package server

import (
	"QiqiLike/conf"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Redis 连接池全局对象
var Cache *redis.Pool

func init() {
	addr := conf.Viper.GetString("redis.addr")
	port := conf.Viper.GetInt("redis.port")
	host := fmt.Sprintf("%s:%d", addr, port)
	// 建立连接池
	Cache = &redis.Pool{
		MaxIdle:     conf.Viper.GetInt("redis.maxIdle"),
		MaxActive:   conf.Viper.GetInt("redis.maxActive"),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			// 选择db
			_, _ = c.Do("SELECT", conf.Viper.GetInt("redis.database"))
			return c, nil
		},
	}
}

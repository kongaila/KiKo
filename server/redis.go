package server

import (
	"KiKo/conf"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// redis初始化
func initRedis() {
	addr := conf.Viper.GetString("redis.addr")
	port := conf.Viper.GetInt("redis.port")
	host := fmt.Sprintf("%s:%d", addr, port)

	var dial = func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", host)
		if err != nil {
			return nil, err
		}
		// 选择db
		_, _ = c.Do("SELECT", conf.Viper.GetInt("redis.database"))
		return c, nil
	}

	// 尝试连接测试
	_, err := dial()
	if err != nil {
		panic(err)
	}

	// 建立连接池
	Cache = &redis.Pool{
		MaxIdle:     conf.Viper.GetInt("redis.maxIdle"),
		MaxActive:   conf.Viper.GetInt("redis.maxActive"),
		IdleTimeout: 180 * time.Second,
		Dial:        dial,
	}
}

package server

import "github.com/garyburd/redigo/redis"

// Redis 连接池全局对象
var Cache *redis.Pool

func init() {
	// redis
	initRedis()
}

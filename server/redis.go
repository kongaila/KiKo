package server

import "github.com/garyburd/redigo/redis"

func GetString(key string) (string, error) {
	r := Cache.Get()
	return redis.String(r.Do("get", key))
}

func SetString(key, value string) (string, error) {
	r := Cache.Get()
	return redis.String(r.Do("set", key, value))
}

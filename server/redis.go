package server

import (
	"github.com/garyburd/redigo/redis"
)

// get
func GetString(key string) (string, error) {
	r := Cache.Get()
	defer r.Close()
	return redis.String(r.Do("get", key))
}

// set
func SetString(key, value string) error {
	r := Cache.Get()
	defer r.Close()
	_, err := r.Do("SET", key, value)
	return err
}

// 验证是否存在
func Exists(key string) (ok bool) {
	r := Cache.Get()
	defer r.Close()
	ok, _ = redis.Bool(r.Do("EXISTS", key))
	return
}

// 删除key
func Del(key string) error {
	r := Cache.Get()
	defer r.Close()
	_, err := redis.Bool(r.Do("DEL", key))
	return err
}

// set ex
func SetEx(key, value string, ex int) error {
	r := Cache.Get()
	defer r.Close()
	_, err := r.Do("SET", key, value, "EX", ex)
	return err
}

// set found ex
func SetFoundEx(key string, ex int) error {
	r := Cache.Get()
	defer r.Close()
	_, err := r.Do("EXPIRE", key, ex)
	return err
}

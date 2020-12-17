package redis

import (
	. "QiqiLike/server"
	"github.com/garyburd/redigo/redis"
)

// redis操作封装工具方法

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

// 验证是否存在并返回剩余时间
func Exists(key string) (ex float64, ok bool) {
	r := Cache.Get()
	defer r.Close()
	ok, _ = redis.Bool(r.Do("EXISTS", key))
	value, _ := redis.Int64(r.Do("TTL", key))
	ex = float64(value)
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
func SetEx(key, value string, ex float64) error {
	r := Cache.Get()
	defer r.Close()
	_, err := r.Do("SET", key, value, "EX", ex)
	return err
}

// set found ex
func SetFoundEx(key string, ex int64) error {
	r := Cache.Get()
	defer r.Close()
	_, err := r.Do("EXPIRE", key, ex)
	return err
}

// 加一
func IncrBy(key string) (int, error) {
	r := Cache.Get()
	defer r.Close()
	return redis.Int(r.Do("incr", key))
}

package church_redis

import (
	"gopkg.in/redis.v5"  // (https://github.com/go-redis/redis)
	"errors"
	"time"
)

const address = "localhost:6379"

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})
}

func Ping() string {
	pong, err := Redis().Ping().Result()
	if err != nil {
		return ""
	}
//	cl := Redis().LRange("testkey", 0, -1)
	return pong
}

func Set(key, value string, expiration time.Duration) error {
	return Redis().Set(key, value, expiration).Err()
}

func Get(key string) (string, error) {
	var err error
	val, err := Redis().Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("Key does not exist")
	} else if err != nil {
		return "", err
	}
	return val, err
}

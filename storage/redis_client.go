package storage

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

const (
	address = "localhost:6379"
	ttl = 0
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	client.Ping().Result()
}

func Get(key string) string {
	val, _ := client.Get(key).Result()
	return val
}

func Set(key string, value string) {
	err := client.Set(key, value, ttl).Err()
	if err != nil {
		panic(err)
	}
}

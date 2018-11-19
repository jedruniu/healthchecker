package healthchecker

import (
	"fmt"

	"github.com/go-redis/redis"
)

type redisBasedHealthCheck struct {
	key         string
	redisClient *redis.Client
}

func NewRedisBased(key string) SingleChecker {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisBasedHealthCheck{key, client}
}

func (hc *redisBasedHealthCheck) SingleCheck() bool {
	val, err := hc.redisClient.Exists(hc.key).Result()
	if err != nil {
		fmt.Printf("could not connect to redis: err %v\n", err)
		return false
	}
	return val == 1
}

package healthchecker

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisBasedHealthCheck struct {
	interval    time.Duration
	key         string
	redisClient *redis.Client
}

func NewRedisBasedHealthChecker(key string, interval time.Duration) HealthChecker {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisBasedHealthCheck{interval, key, client}
}

func (hc *redisBasedHealthCheck) IsHealthy() bool {
	val, err := hc.redisClient.Exists(hc.key).Result()
	if err != nil {
		fmt.Printf("could not connect to redis: err %v\n", err)
		return false
	}
	return val == 1
}

func (hc *redisBasedHealthCheck) GetInterval() time.Duration {
	return hc.interval
}

package chache

import (
	"BloodPressure/pkg/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// 初始化redisClient
func InitRedis(config config.Config) {
	redisCfg := config.RedisConfig
	redisClient = redis.NewClient(&redis.Options{
		DB:           redisCfg.Db,
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() {
	_ = redisClient.Close()
}

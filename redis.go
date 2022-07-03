package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

type IRedisConfig interface {
	GetURL() string
	GetPassword() string
	GetDBIndex() int
}

func InitRedis(conf IRedisConfig) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetURL(),
		Password: conf.GetPassword(),
		DB:       conf.GetDBIndex(),
	})
}

type RedisMessageQueue struct{}

func (r RedisMessageQueue) Push(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := redisClient.RPush(ctx, "mq_"+key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisMessageQueue) PushJSON(key string, value interface{}) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Push(key, string(body))
}

func (r RedisMessageQueue) WorkerHandle(key string, handler func(value string)) {
	for {
		ctx := context.Background()
		result, err := redisClient.BLPop(ctx, 0, key).Result()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			handler(string(result[1]))
		}
	}
}

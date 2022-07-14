package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := redisClient.Ping(c).Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("redis init success")
	}
}

func SetRedisClient(client *redis.Client) {
	redisClient = client
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := redisClient.Ping(c).Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("redis Connect success")
	}
}

type RedisMessageQueue struct {
	Prefix string
}

func (r *RedisMessageQueue) Push(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := redisClient.RPush(ctx, r.Prefix+key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisMessageQueue) PushJSON(key string, value interface{}) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Push(key, string(body))
}

func (r *RedisMessageQueue) WorkerHandle(key string, handler func(value string)) {
	for {
		ctx := context.Background()
		result, err := redisClient.BLPop(ctx, 0, r.Prefix+key).Result()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			handler(string(result[1]))
		}
	}
}

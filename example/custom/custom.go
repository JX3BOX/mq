package main

import (
	"fmt"

	"github.com/JX3BOX/mq"
)

type RedisConfig struct {
	URL      string
	Password string
}

func (r *RedisConfig) GetURL() string {
	return r.URL
}

func (r *RedisConfig) GetPassword() string {
	return r.Password
}
func (r *RedisConfig) GetDBIndex() int {
	return 2
}

func main() {
	var conf = &RedisConfig{
		URL: "127.0.0.1:6379",
	}
	mq.InitRedis(conf)
	var queue = mq.RedisMessageQueue{Prefix: "mq:"}

	queue.WorkerHandle("test", func(value string) {
		fmt.Println("custom", value)
	})
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     conf.GetURL(),
	// 	Password: conf.GetPassword(),
	// 	DB:       conf.GetDBIndex(),
	// })
	// redisClient.Ping(context.Background())
	// for {
	// 	ctx := context.Background()
	// 	result, err := redisClient.BLPop(ctx, 0, "mq_test").Result()
	// 	log.Println(result)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		log.Println(result)
	// 	}
	// }
}

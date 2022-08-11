package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	ctx, cancel := context.WithCancel(context.Background())
	var queue = mq.RedisMessageQueue{Prefix: "mq-dev:", Context: ctx, Cancel: cancel}

	go func() {
		time.Sleep(3 * time.Second)
		queue.Stop()
		log.Println(111)
	}()

	queue.WorkerHandle("test", func(value string) {
		fmt.Println("custom", value)
	})

}

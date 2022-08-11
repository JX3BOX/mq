package main

import (
	"context"
	"fmt"
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

	// go func() {
	// 	queue.Stop()
	// 	log.Println(111)
	// }()

	go queue.WorkerHandle("test", func(value string) {
		fmt.Println("start custom", value)
		time.Sleep(2 * time.Second)
		fmt.Println("end custom", value)
	})
	go queue.WorkerHandle("test1", func(value string) {
		fmt.Println("start custom1", value)
		time.Sleep(2 * time.Second)
		fmt.Println("end custom1", value)
	})
	queue.WorkerHandle("test2", func(value string) {
		fmt.Println("start custom2", value)
		time.Sleep(2 * time.Second)
		fmt.Println("end custom2", value)
	})
	time.Sleep(time.Second)
	queue.Stop()
}

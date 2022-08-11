package main

import (
	"context"
	"fmt"
	"log"

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
	var queue = mq.RedisMessageQueue{Prefix: "mq-dev:", Context: context.Background()}
	// c := time.NewTicker(time.Second)
	// for range c.C {
	// 	v := fmt.Sprintf("productor:%d", time.Now().Unix())
	// 	log.Println(v)
	// 	if err := queue.Push("test", v); err != nil {
	// 		log.Println(err)
	// 	}
	// }
	for i := 0; i < 3; i++ {
		v := fmt.Sprintf("productor:%d", i)
		log.Println(v)
		if err := queue.Push("test", v); err != nil {
			log.Println(err)
		}
		if err := queue.Push("test1", v); err != nil {
			log.Println(err)
		}
		if err := queue.Push("test2", v); err != nil {
			log.Println(err)
		}
	}
}

package main

import (
	"context"
	"log"
	"net/http"
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

	// sigc := make(chan os.Signal, 1)
	// signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	// go func(c chan os.Signal) {
	// 	<-c
	// 	log.Println("prepare close mq")
	// 	os.Exit(0)
	// }(sigc)
	go func() {
		time.Sleep(time.Millisecond * 600)
		log.Println("prepare close mq")
		queue.Stop()
		log.Println("mq has closed")
	}()
	go queue.WorkerHandle("test", func(value string) {
		log.Println("start test:", value)
		time.Sleep(2 * time.Second)
		log.Println("end test:", value)
	})
	go queue.WorkerHandle("test1", func(value string) {
		log.Println("start test1:", value)
		time.Sleep(2 * time.Second)
		log.Println("end test1:", value)
	})
	go queue.WorkerHandle("test2", func(value string) {
		log.Println("start test2:", value)
		time.Sleep(2 * time.Second)
		log.Println("end test2:", value)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":18992", nil)
}

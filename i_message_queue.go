package mq

type IMessageQueue interface {
	Push(key string, value interface{}) error
	PushJSON(key string, value interface{}) error
	WorkerHandle(key string, handler func(value string))
}

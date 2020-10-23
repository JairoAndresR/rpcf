package queue

import "github.com/AgileBits/go-redis-queue/redisqueue"

type Client interface {
	Push(name, content string) error
	GetQueue(name string) *redisqueue.Queue
}

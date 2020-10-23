package queue

import (
	"context"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gomodule/redigo/redis"
	"os"
	"rpcf/core"
	"strconv"
)

func init() {
	err := core.Injector.Provide(newRedisClient)
	core.CheckInjection(err, "newRedisClient")
}

type redisClient struct {
	ctx  context.Context
	conn redis.Conn
}

func newRedisClient() (Client, error) {
	addr := os.Getenv("QUEUE_ADDR")
	password := os.Getenv("QUEUE_PASSWORD")
	db := os.Getenv("QUEUE_DB")
	database, err := strconv.Atoi(db)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	c, err := redis.Dial("tcp", addr, redis.DialDatabase(database), redis.DialPassword(password))
	return &redisClient{conn: c, ctx: ctx}, nil
}

func (r redisClient) GetQueue(name string) *redisqueue.Queue {
	return redisqueue.New(name, r.conn)
}

func (r redisClient) Push(name, content string) error {
	q := r.GetQueue(name)
	_, err := q.Push(content)
	return err
}

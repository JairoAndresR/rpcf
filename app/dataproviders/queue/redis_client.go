package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"rpcf/core"
	"strconv"
)

func init() {
	err := core.Injector.Provide(newRedisClient)
	core.CheckInjection(err, "newRedisClient")
}

type redisClient struct {
	rdb *redis.Client
	ctx context.Context
}

func newRedisClient() (Client, error) {
	addr := os.Getenv("QUEUE_ADDR")
	password := os.Getenv("QUEUE_PASSWORD")
	db := os.Getenv("QUEUE_DB")
	database, err := strconv.Atoi(db)

	if err != nil {
		return nil, err
	}

	opts := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	}
	rdb := redis.NewClient(opts)
	ctx := context.Background()

	return &redisClient{rdb: rdb, ctx: ctx}, nil
}

func (r redisClient) Push(key, content string) error {
	return r.rdb.Set(r.ctx, key, content, 0).Err()
}

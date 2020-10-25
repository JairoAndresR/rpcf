package main

import (
	"fmt"
	"time"

	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gomodule/redigo/redis"
)

func getConnection() (redis.Conn, error) {
	addr := "localhost:6379"
	password := "BUA2WBYgWq9ALAG3"
	database := 0

	c, err := redis.Dial("tcp", addr, redis.DialDatabase(database), redis.DialPassword(password))

	return c, err
}

func getRedisQueue() *redisqueue.Queue {
	c, err := getConnection()

	if err != nil {
		fmt.Println(err)
	} else {
		q := redisqueue.New("collector_queue", c)
		return q
	}
	return nil
}

func main() {

	q := getRedisQueue()
	if q != nil {
		for true {
			job, err := q.Pop()
			if err != nil {
				fmt.Println(err)
			}
			if job != "" {
				fmt.Println(job)
			} else {
				time.Sleep(2 * time.Second)
			}
		}
	}
}

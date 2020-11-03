package main

import (
	"fmt"
	"rpcf/app/dataproviders/queue"
	"time"
)

type productsWorker struct {
	queue queue.Client
}

func newProductsWorker(client queue.Client) *productsWorker {
	return &productsWorker{
		queue: client,
	}
}

func (w *productsWorker) Execute() {
	name := "collector_queue"
	q := w.queue.GetQueue(name)

	if q != nil {
		for true {
			job, err := q.Pop()
			if err != nil {
				panic(err)
			}
			if job != "" {
				fmt.Println(job)
			}
			time.Sleep(2 * time.Second)
		}
	}
}

package main

import (
	"fmt"
	"rpcf/app/dataproviders/queue"
	"rpcf/products/ports"
)

const (
	productsQueue = "collector_queue"
)

type productsWorker struct {
	queue     queue.Client
	collector ports.ProductCollector
}

func newProductsWorker(client queue.Client) *productsWorker {
	collector := NewCollector()
	return &productsWorker{
		queue:     client,
		collector: collector,
	}
}

func (w *productsWorker) Execute() {
	q := w.queue.GetQueue(productsQueue)

	if q == nil {
		return
	}

	for {
		content, err := q.Pop()
		if err != nil {
			panic(err)
		}

		if content == "" {
			continue
		}

		results, errs := w.collector.Parse(content)

		if len(errs) > 0 {
			fmt.Print(errs)
			continue
		}

		fmt.Println(len(results))
	}
}

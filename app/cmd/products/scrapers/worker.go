package main

import (
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
	for true {
		content, err := q.Pop()
		if err != nil {
			panic(err)
		}

		if content != "" {
			w.collector.Parse(content)
		}
	}
}

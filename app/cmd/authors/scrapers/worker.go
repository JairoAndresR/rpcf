package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"rpcf/app/dataproviders/queue"
	"rpcf/products/ports"
)

const (
	authorsQueue = "authors_queue"
)

type authorsWorker struct {
	queue     queue.Client
	collector ports.AuthorsCollector
}

func newAuthorsWorker(client queue.Client) *authorsWorker {
	collector := NewCollector()

	return &authorsWorker{
		queue:     client,
		collector: collector,
	}
}

func (w *authorsWorker) Execute() {
	log.Printf("Worker for authors parsing")

	q := w.queue.GetQueue(authorsQueue)
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

		_, errs := w.collector.Process(content)

		if len(errs) > 0 {
			fmt.Println(errs)
			continue
		}
	}
}

package repositories

import (
	"os"
	"rpcf/app/dataproviders/queue"
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorWriter)
	core.CheckInjection(err, "CollectorWriter")
}

type collectorWriter struct {
	queueName string
	queue     queue.Client
}

func newCollectorWriter(queue queue.Client) ports.CollectorWriter {
	queueName := os.Getenv("COLLECTOR_QUEUE_NAME")
	return &collectorWriter{queue: queue, queueName: queueName}
}

func (c *collectorWriter) Write(payload gruplacs.CollectorPayload) error {
	content, err := payload.JSONString()
	if err != nil {
		return err
	}

	return c.queue.Push(c.queueName, content)
}

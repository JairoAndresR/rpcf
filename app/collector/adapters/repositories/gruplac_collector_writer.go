package repositories

import (
	"os"
	"rpcf/app/dataproviders/queue"
	"rpcf/collector"
	"rpcf/collector/ports"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newCollectorWriter)
	core.CheckInjection(err, "GrupLACCollectorWriter")
}

type GrupLACCollectorWriter struct {
	queueName string
	queue     queue.Client
}

func newCollectorWriter(queue queue.Client) ports.GrupLACCollectorWriter {
	queueName := os.Getenv("COLLECTOR_QUEUE_NAME")
	return &GrupLACCollectorWriter{queue: queue, queueName: queueName}
}

func (c *GrupLACCollectorWriter) Write(payload collector.Payload) error {
	content, err := payload.JSONString()
	if err != nil {
		return err
	}

	return c.queue.Push(c.queueName, content)
}

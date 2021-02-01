package repositories

import (
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
	queue queue.Client
}

func newCollectorWriter(queue queue.Client) ports.GrupLACCollectorWriter {
	return &GrupLACCollectorWriter{queue: queue}
}

func (c *GrupLACCollectorWriter) Write(payload collector.Payload, queueName string) error {
	content, err := payload.JSONString()
	if err != nil {
		return err
	}

	return c.queue.Push(queueName, content)
}

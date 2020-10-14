package repositories

import (
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
	 queue queue.Client
}

func newCollectorWriter(queue queue.Client) ports.CollectorWriter {
	return &collectorWriter{queue: queue}
}

func (c *collectorWriter) Write(payload gruplacs.CollectorPayload) error {
	content, err := payload.JSONString()
	if err != nil {
		return err
	}
	return c.queue.Push(payload.GrupLACCode, content)
}

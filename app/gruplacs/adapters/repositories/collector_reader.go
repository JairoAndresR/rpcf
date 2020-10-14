package repositories

import (
	"rpcf/core"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorReader)
	core.CheckInjection(err, "CollectorReader")
}

type collectorReader struct{}

func newCollectorReader() ports.CollectorReader {
	return &collectorReader{}
}

func (c *collectorReader) GetContent(URL string) string {
	return ""
}

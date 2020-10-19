package repositories

import (
	"github.com/udistritali3plus/collector"
	"rpcf/core"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorReader)
	core.CheckInjection(err, "CollectorReader")
}

type collectorReader struct {
	collector collector.Collector
}

func newCollectorReader() ports.CollectorReader {
	c := collector.NewCollector()

	return &collectorReader{
		collector: c,
	}
}

func (c *collectorReader) GetContent(url string) (string, error) {
	return c.collector.GetContent(url)
}

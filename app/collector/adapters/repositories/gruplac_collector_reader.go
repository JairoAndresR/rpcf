package repositories

import (
	"github.com/udistritali3plus/collector"
	"rpcf/collector/ports"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newCollectorReader)
	core.CheckInjection(err, "GRUPLACCollectorReader")
}

type GrupLACCollectorReader struct {
	collector collector.Collector
}

func newCollectorReader() ports.GRUPLACCollectorReader {
	c := collector.NewCollector()

	return &GrupLACCollectorReader{
		collector: c,
	}
}

func (c *GrupLACCollectorReader) GetContent(url string) (string, error) {
	return c.collector.GetContent(url)
}

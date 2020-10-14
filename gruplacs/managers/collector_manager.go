package managers

import (
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorManager)
	core.CheckInjection(err, "CollectorManager")
}

type collectorManager struct {
	grupLACReader   ports.GrupLACReader
	collectorReader ports.CollectorReader
	collectorWriter ports.CollectorWriter
}

func newCollectorManager(gr ports.GrupLACReader, cr ports.CollectorReader, cw ports.CollectorWriter) ports.CollectorManager {
	return &collectorManager{
		grupLACReader:   gr,
		collectorReader: cr,
		collectorWriter: cw,
	}
}

func (c *collectorManager) Collect(researchGroup string) {

	groups := c.grupLACReader.GetAll()

	for _, g := range groups {
		content := c.collectorReader.GetContent(g.URL)
		code := g.GetCode()
		payload := gruplacs.NewCollectorPayload(content, code)
		c.collectorWriter.Write(*payload)
	}
}

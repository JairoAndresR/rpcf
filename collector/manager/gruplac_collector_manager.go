package manager

import (
	"rpcf/collector/ports"
	"rpcf/core"
	gruplacPorts "rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorManager)
	core.CheckInjection(err, "newCollectorManager")
}

const (
	productsQueue = "products_queue"
)

type GrupLACCollectorManager struct {
	*collectorBase
}

func newCollectorManager(
	gr gruplacPorts.GruplacDefinitionReader,
	cr ports.GRUPLACCollectorReader,
	cw ports.GrupLACCollectorWriter) ports.GrupLACCollectorManager {
	base := &collectorBase{
		grupLACReader:   gr,
		collectorReader: cr,
		collectorWriter: cw,
	}

	return &GrupLACCollectorManager{
		base,
	}
}

func (c *GrupLACCollectorManager) CollectAll() error {
	return c.Collect(productsQueue)
}

package managers

import (
	"rpcf/collector/ports"
	"rpcf/core"
	gruplacPorts "rpcf/gruplacs/ports"
)

const (
	authorsQueue = "authors_queue"
)

func init() {
	err := core.Injector.Provide(newAuthorsCollectorManager)
	core.CheckInjection(err, "newAuthorsCollectorManager")
}

type authorsCollectorManager struct {
	*collectorBase
}

func newAuthorsCollectorManager(
	gr gruplacPorts.GruplacDefinitionReader,
	cr ports.GRUPLACCollectorReader,
	cw ports.GrupLACCollectorWriter,
) ports.AuthorsCollectorManager {

	base := &collectorBase{
		grupLACReader:   gr,
		collectorReader: cr,
		collectorWriter: cw,
	}

	return &authorsCollectorManager{
		base,
	}
}
func (m *authorsCollectorManager) CollectAll() error {
	return m.Collect(authorsQueue)
}

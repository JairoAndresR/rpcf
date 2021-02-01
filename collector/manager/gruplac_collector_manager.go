package manager

import (
	"errors"
	"rpcf/collector"
	"rpcf/collector/ports"
	"rpcf/core"
	"rpcf/gruplacs/managers"
	gruplacPorts "rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newCollectorManager)
	core.CheckInjection(err, "GrupLACCollectorManager")
}

type GrupLACCollectorManager struct {
	grupLACReader   gruplacPorts.GruplacDefinitionReader
	collectorReader ports.GRUPLACCollectorReader
	collectorWriter ports.GrupLACCollectorWriter
}

func newCollectorManager(gr gruplacPorts.GruplacDefinitionReader, cr ports.GRUPLACCollectorReader, cw ports.GrupLACCollectorWriter) ports.GrupLACCollectorManager {
	return &GrupLACCollectorManager{
		grupLACReader:   gr,
		collectorReader: cr,
		collectorWriter: cw,
	}
}

func (c *GrupLACCollectorManager) CollectAll() error {

	groups, err := c.grupLACReader.GetAll()

	if err != nil {
		return errors.New(managers.NotPossibleRetrieveGrupLACsError)
	}

	for _, g := range groups {
		content, err := c.collectorReader.GetContent(g.URL)
		if err != nil {
			// TODO log it and marks as not readable
			return err
		}
		code := g.GetCode()
		payload := collector.NewCollectorPayload(content, code, g.Name)
		c.collectorWriter.Write(*payload)
	}

	return nil
}

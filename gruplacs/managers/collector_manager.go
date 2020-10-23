package managers

import (
	"errors"
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

func (c *collectorManager) CollectAll() error {

	groups, err := c.grupLACReader.GetAll()

	if err != nil {
		return errors.New(NotPossibleRetrieveGrupLACsError)
	}

	for _, g := range groups {
		content, err := c.collectorReader.GetContent(g.URL)
		if err != nil {
			// TODO log it and marks as not readable
			return err
		}
		code := g.GetCode()
		payload := gruplacs.NewCollectorPayload(content, code)
		c.collectorWriter.Write(*payload)
	}
	return nil
}

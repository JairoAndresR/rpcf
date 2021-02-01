package manager

import (
	"errors"
	"rpcf/collector"
	"rpcf/collector/ports"
	"rpcf/gruplacs/managers"
	gruplacPorts "rpcf/gruplacs/ports"
)

type collectorBase struct {
	grupLACReader   gruplacPorts.GruplacDefinitionReader
	collectorReader ports.GRUPLACCollectorReader
	collectorWriter ports.GrupLACCollectorWriter
}

func (c *collectorBase) Collect(productsQueue string) error {

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
		c.collectorWriter.Write(*payload, productsQueue)
	}

	return nil
}

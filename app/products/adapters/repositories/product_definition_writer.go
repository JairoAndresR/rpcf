package repositories

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductDefinitionWriter)
	core.CheckInjection(err, "productDefinitionWriter")
}

type productDefinitionWriter struct {
}

func newProductDefinitionWriter() ports.ProductDefinitionWriter {
	return &productDefinitionWriter{}
}

func (p2 productDefinitionWriter) Create(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	panic("implement me")
}

func (p2 productDefinitionWriter) Update(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	panic("implement me")
}

func (p2 productDefinitionWriter) Delete(id string) error {
	panic("implement me")
}

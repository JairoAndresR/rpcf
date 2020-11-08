package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductDefinitionManager)
	core.CheckInjection(err, "newProductDefinitionManager")
}

type productDefinitionManager struct {
	writer ports.ProductDefinitionWriter
	reader ports.ProductDefinitionReader
}

func newProductDefinitionManager(
	writer ports.ProductDefinitionWriter,
	reader ports.ProductDefinitionReader) ports.ProductDefinitionManager {

	return &productDefinitionManager{
		writer,
		reader}
}

func (m *productDefinitionManager) Create(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	return m.writer.Create(p)
}

func (m *productDefinitionManager) Update(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	return m.writer.Update(p)
}

func (m *productDefinitionManager) Delete(id string) error {
	return m.writer.Delete(id)
}

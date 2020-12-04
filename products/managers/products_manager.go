package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductsManager)
	core.CheckInjection(err, "newProductsManager")
}

type productsManager struct {
	reader ports.ProductReader
}

func newProductsManager(reader ports.ProductReader) ports.ProductsManager {
	return &productsManager{
		reader: reader,
	}
}

func (m *productsManager) GetAll(filters map[string]string) ([]*products.Product, error) {
	return m.reader.GetAll(filters)
}

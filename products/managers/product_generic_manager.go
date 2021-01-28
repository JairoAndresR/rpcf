package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductGenericManager)
	core.CheckInjection(err, "newProductGenericManager")
}

type productGenericManager struct {
	writer ports.GenericProductWriter
}

func newProductGenericManager(writer ports.GenericProductWriter) ports.ProductGenericManager {
	return &productGenericManager{writer: writer}
}

func (m productGenericManager) Write(product *products.ProductResult) (ports.GenericProduct, error) {
	return m.writer.WriteMap(product)
}

package managers

import (
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductManager)
	core.CheckInjection(err, "newProductManager")
}

type productManager struct {
	writer ports.ProductGenericWriter
}

func newProductManager(writer ports.ProductGenericWriter) ports.ProductManager {
	return &productManager{writer: writer}
}

func (m productManager) Write(product map[string]string, name string) error {
	return m.writer.WriteMap(product, name)
}

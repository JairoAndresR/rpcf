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
	writer ports.ProductWriter
}

func newProductManager(writer ports.ProductWriter) ports.ProductManager {
	return &productManager{writer: writer}
}

func (m productManager) Write(product map[string]string, name string) error {
	return m.writer.Write(product, name)
}
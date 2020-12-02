package managers

import (
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductGenericManager)
	core.CheckInjection(err, "newProductGenericManager")
}

type productGenericManager struct {
	writer ports.ProductGenericWriter
}

func newProductGenericManager(writer ports.ProductGenericWriter) ports.ProductGenericManager {
	return &productGenericManager{writer: writer}
}

func (m productGenericManager) Write(product map[string]string, name string) error {
	return m.writer.WriteMap(product, name)
}

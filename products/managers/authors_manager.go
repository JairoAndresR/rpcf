package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorsManager)
	core.CheckInjection(err, "newAuthorsManager")
}

type authorsManager struct {
	reader ports.AuthorsReader
}

func newAuthorsManager(ar ports.AuthorsReader) ports.AuthorsManager {
	return &authorsManager{
		reader: ar,
	}
}

func (m *authorsManager) GetAll() ([]*products.Author, error) {
	return m.reader.GetAll()
}

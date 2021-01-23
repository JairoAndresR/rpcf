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

func (m *authorsManager) GetAll(groupCode string) ([]*products.Author, error) {
	if groupCode == "" {
		return m.reader.GetAll()
	}

	return m.reader.GetAllByGroup(groupCode)
}

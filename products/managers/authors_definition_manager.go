package managers

import (
	"errors"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorDefinitionManager)
	core.CheckInjection(err, "newAuthorDefinitionManager")
}

type authorDefinitionManager struct {
	reader ports.AuthorsDefinitionReader
	writer ports.AuthorsDefinitionWriter
}

func newAuthorDefinitionManager(
	r ports.AuthorsDefinitionReader,
	w ports.AuthorsDefinitionWriter,

) ports.AuthorsDefinitionManager {
	return &authorDefinitionManager{
		reader: r,
		writer: w,
	}
}

func (m authorDefinitionManager) Create(d *products.AuthorDefinition) (*products.AuthorDefinition, error) {
	d, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	if d != nil {
		return nil, errors.New(authorDefinitionAlreadyExist)
	}

	return m.writer.Create(d)
}

func (m authorDefinitionManager) GetAll() (*products.AuthorDefinition, error) {
	return m.reader.GetAuthorDefinition()
}

package repositories

import (
	"gorm.io/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorDefinitionWriter)
	core.CheckInjection(err, "newAuthorDefinitionWriter")
}

type authorDefinitionWriter struct {
	db *gorm.DB
}

func newAuthorDefinitionWriter(con sql.Connection) ports.AuthorsDefinitionWriter {
	return &authorDefinitionWriter{db: con.GetDatabase()}
}

func (w *authorDefinitionWriter) Create(d *products.AuthorDefinition) (*products.AuthorDefinition, error) {
	definition := entities.NewAuthorDefinition(d)
	err := w.db.Create(definition).Error
	return definition.GetDomainReference(), err
}

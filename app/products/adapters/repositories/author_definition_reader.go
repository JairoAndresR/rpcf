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
	err := core.Injector.Provide(newAuthorDefinitionReader)
	core.CheckInjection(err, "newAuthorDefinitionReader")
}

type AuthorDefinitionReader struct {
	db *gorm.DB
}

func newAuthorDefinitionReader(con sql.Connection) ports.AuthorsDefinitionReader {
	db := con.GetDatabase()
	return &AuthorDefinitionReader{db: db}
}

func (r *AuthorDefinitionReader) GetAuthorDefinition() (*products.AuthorDefinition, error) {
	var definition entities.AuthorDefinition
	err := r.db.First(&definition).Error
	if err != nil {
		return nil, err
	}
	return definition.GetDomainReference(), nil
}

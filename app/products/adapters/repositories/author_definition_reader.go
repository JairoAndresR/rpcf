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

func newAuthorDefinitionReader(con sql.Connection) ports.AuthorDefinitionReader {
	db := con.GetDatabase()
	return &AuthorDefinitionReader{db: db}
}

func (r *AuthorDefinitionReader) GetAuthorDefinition() (*products.AuthorDefinition, error) {
	var definition *entities.AuthorDefinition
	err := r.db.Find(&definition).Error
	return definition.GetDomainReference(), err
}

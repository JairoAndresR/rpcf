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
	err := core.Injector.Provide(newProductDefinitionReader)
	core.CheckInjection(err, "newProductDefinitionReader")
}

type productDefinitionReader struct {
	db *gorm.DB
}

func newProductDefinitionReader(con sql.Connection) ports.ProductDefinitionReader {
	db := con.GetDatabase()
	return &productDefinitionReader{db: db}
}

func (r *productDefinitionReader) GetAll() ([]*products.ProductDefinition, error) {
	var definitions []*entities.ProductDefinition
	err := r.db.Find(&definitions).Error
	return entities.MapProductDefinitionListToDomain(definitions), err
}

func (r *productDefinitionReader) GetById(id string) (*products.ProductDefinition, error) {
	var definition entities.ProductDefinition
	err := r.db.Where("id = ?", id).First(&definition).Error
	return definition.GetDomainReference(), err
}

func (r *productDefinitionReader) GetByName(name string) (*products.ProductDefinition, error) {
	var definition entities.ProductDefinition
	err := r.db.First(&definition).Where("name = ?", name).Error
	return definition.GetDomainReference(), err
}

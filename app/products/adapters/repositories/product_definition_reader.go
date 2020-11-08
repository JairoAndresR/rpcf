package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductReader)
	core.CheckInjection(err, "newProductReader")
}

type productDefinitionReader struct {
	db *gorm.DB
}

func newProductReader(con sql.Connection) ports.ProductDefinitionReader {
	db := con.GetDatabase()
	return &productDefinitionReader{db: db}
}

func (r *productDefinitionReader) GetAll() ([]products.ProductDefinition, error) {
	var definitions []*entities.ProductDefinition
	err := r.db.Find(&definitions).Error
	return entities.MapListToDomain(definitions), err
}

func (r *productDefinitionReader) GetById(id string) (*products.ProductDefinition, error) {
	var definition entities.ProductDefinition
	err := r.db.First(&definition).Where("id = ?", id).Error
	return definition.GetDomainReference(), err
}

func (r *productDefinitionReader) GetByName(name string) (*products.ProductDefinition, error) {
	var definition entities.ProductDefinition
	err := r.db.First(&definition).Where("name = ?", name).Error
	return definition.GetDomainReference(), err
}

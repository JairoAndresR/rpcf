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
	err := core.Injector.Provide(newProductDefinitionWriter)
	core.CheckInjection(err, "productDefinitionWriter")
}

type productDefinitionWriter struct {
	db *gorm.DB
}

func newProductDefinitionWriter(conn sql.Connection) ports.ProductDefinitionWriter {
	db := conn.GetDatabase()
	return &productDefinitionWriter{db: db}
}

func (w productDefinitionWriter) Create(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	definition := entities.NewProductDefinition(p)
	err := w.db.Create(definition).Error
	return definition.GetDomainReference(), err
}

func (w productDefinitionWriter) Update(p *products.ProductDefinition) (*products.ProductDefinition, error) {
	definition := entities.NewProductDefinition(p)
	err := w.db.Model(&entities.ProductDefinition{}).Updates(definition).Error
	return definition.GetDomainReference(), err
}

func (w productDefinitionWriter) Delete(id string) error {
	return w.db.Delete(entities.ProductDefinition{}, "id = ?", id).Error
}

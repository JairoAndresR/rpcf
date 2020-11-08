package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products"
	ports2 "rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductReader)
	core.CheckInjection(err, "newProductReader")
}

type productDefinitionReader struct {
	db *gorm.DB
}

func newProductReader(con sql.Connection) ports2.ProductDefinitionReader {
	db := con.GetDatabase()
	return &productDefinitionReader{db: db}
}

func (r *productDefinitionReader) GetAll() ([]products.ProductDefinition, error) {
	var definitions []*entities.ProductDefinition
	err := r.db.Find(&definitions).Error
	return entities.MapListToDomain(definitions), err
}

func (r *productDefinitionReader) GetById(id string) (*products.ProductDefinition, error) {
	panic("implement me")
}

func (r *productDefinitionReader) GetByName(name string) (*products.ProductDefinition, error) {
	panic("implement me")
}

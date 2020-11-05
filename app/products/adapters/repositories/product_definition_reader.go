package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/app/products/ports"
	"rpcf/core"
	"rpcf/products"
)

func init(){
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

func (r *productDefinitionReader) GetAll() ([]products.ProductDefinition, error){
	var definitions []*entities.ProductDefinition
	err:= r.db.Find(&definitions).Error
	return entities.MapListToDomain(definitions), err
}
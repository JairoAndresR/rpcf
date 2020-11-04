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

type productReader struct {
	db *gorm.DB
}

func newProductReader(con sql.Connection) ports.ProductReader{
	db := con.GetDatabase()
	return &productReader{db: db}
}

func (g *productReader) GetAll() ([]products.Product, error){
	products := make([]*entities.Product, 0)

	product:= &entities.Product{
		Name: "Product 1",
		Definition: "Definition product 1",
	}
	products =  append(products, product)

	return entities.MapListToDomain(products), nil
}
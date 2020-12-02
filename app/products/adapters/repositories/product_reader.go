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

type productReader struct {
	db *gorm.DB
}

func newProductReader(conn sql.Connection) ports.ProductReader {
	return &productReader{db: conn.GetDatabase()}
}

func (r *productReader) GetAll() ([]*products.Product, error) {
	var list []*entities.Product
	err := r.db.Find(&list).Error

	if err != nil {
		return []*products.Product{}, err
	}

	products := entities.MapListToDomain(list)
	return products, nil
}

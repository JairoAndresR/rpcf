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
	err := core.Injector.Provide(newProductWriter)
	core.CheckInjection(err, "newProductWriter")
}

type productWriter struct {
	db *gorm.DB
}

func newProductWriter(conn sql.Connection) ports.ProductWriter {
	return &productWriter{db: conn.GetDatabase()}
}

func (w *productWriter) Create(p *products.Product) (*products.Product, error) {
	product := entities.NewProduct(p)
	err := w.db.Save(product).Error
	return product.ToDomain(), err
}

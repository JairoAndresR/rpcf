package repositories

import (
	"fmt"
	"gorm.io/gorm"
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

func (r *productReader) GetAll(filters map[string]string) ([]*products.Product, error) {
	var list []*entities.Product
	tx := r.db.Model(entities.Product{}).Preload("Authors")
	researcher := filters["researcher_id"]
	
	if researcher != "" {
		delete(filters, "researcher_id")
		tx.Where("id IN (SELECT product_id from rpcf.authors_products WHERE author_id = ?)", researcher)
	}
	
	for column, value := range filters {
		q := fmt.Sprintf(`%s LIKE ?`, column)
		tx.Where(q, fmt.Sprintf("%%%s%%", value))
	}
	err := tx.Find(&list).Error
	if err != nil {
		return nil, err
	}
	products := entities.MapListToDomain(list)
	return products, nil
}

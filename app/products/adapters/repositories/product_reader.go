package repositories

import (
	"errors"
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
	tx := r.db.Model(entities.Product{})
	for column, value := range filters {
		q := fmt.Sprintf(`%s LIKE ?`, column)
		tx.Where(q, fmt.Sprintf("%%%s%%", value))
	}
	err := tx.Find(&list).Error
	if err != nil {
		return nil, errors.New("find error")
	}
	products := entities.MapListToDomain(list)
	return products, nil
}

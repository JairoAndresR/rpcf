package repositories

import (
	"gorm.io/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/clauses"
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

	predicate := clauses.NewProductFilterPredicate(r.db)
	values, sql := predicate.Build(filters)
	err := r.db.Raw(sql, values).Scan(&list).Error

	if err != nil {
		return []*products.Product{}, err
	}

	products := entities.MapListToDomain(list)
	return products, nil
}

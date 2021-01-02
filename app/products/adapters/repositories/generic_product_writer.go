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
	err := core.Injector.Provide(newGenericProductWriter)
	core.CheckInjection(err, "newGenericProductWriter")
}

type genericProductWriter struct {
	db      *gorm.DB
	builder ports.ProductsBuilder
}

func newGenericProductWriter(conn sql.Connection, b ports.ProductsBuilder) ports.GenericProductWriter {
	db := conn.GetDatabase()
	return &genericProductWriter{db: db, builder: b}
}

func (w *genericProductWriter) WriteGeneric(product map[string]string, name string) (*products.Product, error) {
	p, err := w.WriteMap(product, name)
	if err != nil {
		return nil, err
	}

	p, err = w.Write(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (w *genericProductWriter) Write(product *products.Product) (*products.Product, error) {
	p := entities.NewProduct(product)
	err := w.db.Model(p).Create(p).Error
	if err != nil {
		return nil, err
	}

	return p.ToDomain(), nil
}

func (w *genericProductWriter) WriteMap(product map[string]string, name string) (*products.Product, error) {
	p, err := w.builder.Build(product, name)

	if err != nil {
		return nil, err
	}

	err = w.db.Table(name).Create(p).Error
	if err != nil {
		return nil, err
	}
	generic := ports.NewGenericProduct(p, name)
	return generic, err
}

func (w *genericProductWriter) WriteGenerics(ps []map[string]string, name string) ([]*products.Product, []error) {
	errs := make([]error, 0)
	results := make([]*products.Product, 0)

	for _, p := range ps {
		r, err := w.WriteGeneric(p, name)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		results = append(results, r)
	}
	return results, errs
}

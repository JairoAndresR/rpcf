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

func (w *genericProductWriter) WriteGeneric(product *products.ProductResult) (*products.Product, error) {
	p, err := w.WriteMap(product)
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

func (w *genericProductWriter) WriteMap(product *products.ProductResult) (*products.Product, error) {
	p, err := w.builder.Build(product.Fields, product.Name)

	if err != nil {
		return nil, err
	}

	err = w.db.Table(product.Name).Create(p).Error
	if err != nil {
		return nil, err
	}
	generic := ports.NewGenericProduct(p, product.GrupLACCode, product.GrupLACCode, product.Name)
	return generic, err
}

func (w *genericProductWriter) WriteGenerics(results []*products.ProductResult) ([]*products.Product, []error) {
	errs := make([]error, 0)
	res := make([]*products.Product, 0)

	for _, p := range results {
		r, err := w.WriteGeneric(p)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		res = append(res, r)
	}
	return res, errs
}

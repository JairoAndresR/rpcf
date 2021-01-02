package repositories

import (
	"gorm.io/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
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

func (w *genericProductWriter) WriteMap(product map[string]string, name string) (ports.GenericProduct, error) {
	p, err := w.builder.Build(product, name)

	if err != nil {
		return nil, err
	}

	err = w.db.Table(name).Create(p).Error
	if err != nil {
		return nil, err
	}
	generic := ports.NewGenericProduct(p)
	return generic, err
}

func (w *genericProductWriter) WriteMaps(products []map[string]string, name string) ([]ports.GenericProduct, []error) {
	errs := make([]error, 0)
	results := make([]ports.GenericProduct, 0)

	for _, p := range products {
		r, err := w.WriteMap(p, name)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		results = append(results, r)
	}
	return results, errs
}

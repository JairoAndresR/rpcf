package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductWriter)
	core.CheckInjection(err, "newProductWriter")
}

type productWriter struct {
	db      *gorm.DB
	builder ports.ProductsBuilder
}

func newProductWriter(conn sql.Connection, b ports.ProductsBuilder) ports.ProductWriter {
	db := conn.GetDatabase()
	return &productWriter{db: db, builder: b}
}

func (w *productWriter) WriteMap(product map[string]string, name string) error {
	p, err := w.builder.Build(product, name)

	if err != nil {
		return err
	}

	err = w.db.Table(name).Create(p).Error
	return err
}

func (w *productWriter) WriteMaps(products []map[string]string, name string) []error {
	errs := make([]error, 0)

	for _, p := range products {
		err := w.WriteMap(p, name)

		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

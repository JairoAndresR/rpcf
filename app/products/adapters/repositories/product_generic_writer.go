package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductGenericWriter)
	core.CheckInjection(err, "newProductGenericWriter")
}

type productGenericWriter struct {
	db      *gorm.DB
	builder ports.ProductsBuilder
}

func newProductGenericWriter(conn sql.Connection, b ports.ProductsBuilder) ports.ProductGenericWriter {
	db := conn.GetDatabase()
	return &productGenericWriter{db: db, builder: b}
}

func (w *productGenericWriter) WriteMap(product map[string]string, name string) error {
	p, err := w.builder.Build(product, name)

	if err != nil {
		return err
	}

	err = w.db.Table(name).Create(p).Error
	return err
}

func (w *productGenericWriter) WriteMaps(products []map[string]string, name string) []error {
	errs := make([]error, 0)

	for _, p := range products {
		err := w.WriteMap(p, name)

		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

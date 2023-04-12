package managers

import (
	"fmt"
	"rpcf/app/dataproviders/sql"
	"rpcf/collector/ports"
	"rpcf/core"
	gruplacPorts "rpcf/gruplacs/ports"

	"gorm.io/gorm"
)

const (
	authorsQueue = "authors_queue"
)

func init() {
	err := core.Injector.Provide(newAuthorsCollectorManager)
	core.CheckInjection(err, "newAuthorsCollectorManager")
}

type authorsCollectorManager struct {
	db   *gorm.DB
	base *collectorBase
}

func newAuthorsCollectorManager(
	gr gruplacPorts.GruplacDefinitionReader,
	cr ports.GRUPLACCollectorReader,
	cw ports.GrupLACCollectorWriter,
	conn sql.Connection,
) ports.AuthorsCollectorManager {

	base := &collectorBase{
		grupLACReader:   gr,
		collectorReader: cr,
		collectorWriter: cw,
	}

	return &authorsCollectorManager{
		db:   conn.GetDatabase(),
		base: base,
	}
}
func (m *authorsCollectorManager) CollectAll() error {
	m.ClearDB()
	if err := m.base.Collect(authorsQueue); err != nil {
		fmt.Println("CollectAll %s", err)
		return err
	}
	return nil
}

func (w *authorsCollectorManager) ClearDB() {
	w.db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	w.db.Exec("TRUNCATE TABLE authors_products;")
	w.db.Exec("TRUNCATE TABLE authors_gruplacs;")
	w.db.Exec("TRUNCATE TABLE products;")
	w.db.Exec("TRUNCATE TABLE authors;")
	w.db.Exec("TRUNCATE TABLE articles;")
	w.db.Exec("TRUNCATE TABLE books;")
	w.db.Exec("TRUNCATE TABLE companies;")
	w.db.Exec("TRUNCATE TABLE doctoral_theses;")
	w.db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}

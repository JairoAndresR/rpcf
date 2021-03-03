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
	err := core.Injector.Provide(newGrupLACWriter)
	core.CheckInjection(err, "newGrupLACWriter")
}

type grupLACWriter struct {
	db *gorm.DB
}

func newGrupLACWriter(conn sql.Connection) ports.GrupLACWriter {
	return &grupLACWriter{db: conn.GetDatabase()}
}

func (w *grupLACWriter) Create(g *products.GrupLAC) (*products.GrupLAC, error) {
	gruplac := entities.NewGrupLAC(g)
	err := w.db.Create(&gruplac).Error
	return gruplac.ToDomain(), err
}

func (w *grupLACWriter) Update(g *products.GrupLAC) (*products.GrupLAC, error) {
	gruplac := entities.NewGrupLAC(g)
	err := w.db.Save(&gruplac).Error
	return gruplac.ToDomain(), err
}

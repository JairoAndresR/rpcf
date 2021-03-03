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
	err := core.Injector.Provide(newGrupLACReader)
	core.CheckInjection(err, "newGrupLACReader")
}

type grupLACReader struct {
	db *gorm.DB
}

func newGrupLACReader(conn sql.Connection) ports.GrupLACReader {
	return &grupLACReader{db: conn.GetDatabase()}
}

func (r *grupLACReader) GetAll() ([]*products.GrupLAC, error) {
	var gruplacs []*entities.GrupLAC
	err := r.db.Find(&gruplacs).Error
	return entities.MapGrupLACsToLisDomain(gruplacs), err
}

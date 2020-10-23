package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/gruplacs/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newGrupLACReader)
	core.CheckInjection(err, "newGrupLACReader")
}

type grupLACReader struct {
	db *gorm.DB
}

func newGrupLACReader(con sql.Connection) ports.GrupLACReader {
	db := con.GetDatabase()
	return &grupLACReader{db: db}
}

func (g *grupLACReader) GetAll() ([]gruplacs.GrupLAC, error) {
	groupLacs := make([]*entities.GrupLAC, 0)

	group := &entities.GrupLAC{
		Name: "jira",
		URL:  "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394",
	}

	groupLacs = append(groupLacs, group)
	/*var groupLacs []*entities.GrupLAC

	err := g.db.Find(&groupLacs).Error

	if err != nil {
		return []gruplacs.GrupLAC{}, err
	}*/
	return entities.MapListToDomain(groupLacs), nil
}

package repositories

import (
	"gorm.io/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/gruplacs/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newGruplacReader)
	core.CheckInjection(err, "newGruplacReader")
}

type gruplacDefinitionReader struct {
	db *gorm.DB
}

func newGruplacReader(con sql.Connection) ports.GruplacDefinitionReader {
	db := con.GetDatabase()
	return &gruplacDefinitionReader{db: db}
}

func (g *gruplacDefinitionReader) GetAll() ([]*gruplacs.GruplacDefinition, error) {

	var definitions []*entities.GruplacDefinition
	err := g.db.Find(&definitions).Error
	return entities.MapGruplacDefinitionListToDomain(definitions), err
}

func (g *gruplacDefinitionReader) GetById(id string) (*gruplacs.GruplacDefinition, error) {
	var definition entities.GruplacDefinition
	err := g.db.Where("id = ?", id).First(&definition).Error
	return definition.GetDomainReference(), err
}

func (g *gruplacDefinitionReader) GetByName(name string) (*gruplacs.GruplacDefinition, error) {
	var definition entities.GruplacDefinition
	err := g.db.First(&definition).Where("name = ?", name).Error
	return definition.GetDomainReference(), err
}

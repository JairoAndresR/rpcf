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
	err := core.Injector.Provide(newGruplaDefinitionWriter)
	core.CheckInjection(err, "gruplacDefinitionWriter")
}

type gruplacDefinitionWriter struct {
	db *gorm.DB
}

func newGruplaDefinitionWriter(conn sql.Connection) ports.GruplacDefinitionWriter {
	db := conn.GetDatabase()
	return &gruplacDefinitionWriter{db: db}
}

func (w gruplacDefinitionWriter) Create(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	definition := entities.NewGruplacDefinition(g)
	err := w.db.Create(definition).Error
	return definition.GetDomainReference(), err
}

func (w gruplacDefinitionWriter) Update(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	definition := entities.NewGruplacDefinition(g)
	err := w.db.Model(&entities.GruplacDefinition{}).Updates(definition).Error
	return definition.GetDomainReference(), err
}

func (w gruplacDefinitionWriter) Delete(id string) error {
	return w.db.Delete(entities.GruplacDefinition{}, "id = ?", id).Error
}

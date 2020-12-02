package managers

import (
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newGruplacDefinitionManager)
	core.CheckInjection(err, "newGruplacDefinitionManager")
}

type gruplacDefinitionManager struct {
	writer ports.GruplacDefinitionWriter
	reader ports.GruplacDefinitionReader
}

func newGruplacDefinitionManager(
	writer ports.GruplacDefinitionWriter,
	reader ports.GruplacDefinitionReader) ports.GruplacDefinitionManager {

	return &gruplacDefinitionManager{
		writer,
		reader,
	}
}

func (m *gruplacDefinitionManager) Create(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	return m.writer.Create(g)
}

func (m *gruplacDefinitionManager) Update(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	return m.writer.Update(g)
}

func (m *gruplacDefinitionManager) Delete(id string) error {
	return m.writer.Delete(id)
}

func (m *gruplacDefinitionManager) GetAll() ([]*gruplacs.GruplacDefinition, error) {
	return m.reader.GetAll()
}

func (m *gruplacDefinitionManager) GetById(id string) (*gruplacs.GruplacDefinition, error) {
	return m.reader.GetById(id)
}

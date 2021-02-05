package managers

import (
	"rpcf/core"
	"rpcf/gruplacs"
	"rpcf/gruplacs/ports"
	"rpcf/products"
	productPorts "rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newGruplacDefinitionManager)
	core.CheckInjection(err, "newGruplacDefinitionManager")
}

type gruplacDefinitionManager struct {
	writer        ports.GruplacDefinitionWriter
	reader        ports.GruplacDefinitionReader
	gruplacWriter productPorts.GrupLACWriter
}

func newGruplacDefinitionManager(
	writer ports.GruplacDefinitionWriter,
	reader ports.GruplacDefinitionReader,
	gr productPorts.GrupLACWriter,
) ports.GruplacDefinitionManager {

	return &gruplacDefinitionManager{
		writer:        writer,
		reader:        reader,
		gruplacWriter: gr,
	}
}

func (m *gruplacDefinitionManager) Create(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	def, err := m.writer.Create(g)
	m.gruplacWriter.Create(products.NewGrupLAC(def.Code, def.Name))
	return def, err
}

func (m *gruplacDefinitionManager) Update(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error) {
	def, err := m.writer.Update(g)
	m.gruplacWriter.Update(products.NewGrupLAC(def.Code, def.Name))
	return def, err
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

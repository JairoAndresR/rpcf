package ports

import "rpcf/gruplacs"

type GruplacDefinitionManager interface {
	Create(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error)

	Update(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error)

	Delete(id string) error

	GetAll() ([]*gruplacs.GruplacDefinition, error)

	GetById(id string) (*gruplacs.GruplacDefinition, error)
}
